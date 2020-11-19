package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	et "github.com/jackdon/cloudmusic/entities"
)

var client *http.Client

func chooseUserAgent(ua string) string {
	var userAgentList = map[string][]string{
		"mobile": {
			// iOS 13.5.1 14.0 beta with safari
			"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.",
			// iOS with qq micromsg
			"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML like Gecko) Mobile/14A456 QQ/6.5.7.408 V1_IPH_SQ_6.5.7_1_APP_A Pixel/750 Core/UIWebView NetType/4G Mem/103",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.15(0x17000f27) NetType/WIFI Language/zh",
			// Android -> Huawei Xiaomi
			"Mozilla/5.0 (Linux; Android 9; PCT-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.311 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; U; Android 9; zh-cn; Redmi Note 8 Build/PKQ1.190616.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.141 Mobile Safari/537.36 XiaoMi/MiuiBrowser/12.5.22",
			// Android + qq micromsg
			"Mozilla/5.0 (Linux; Android 10; YAL-AL00 Build/HUAWEIYAL-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.62 XWEB/2581 MMWEBSDK/200801 Mobile Safari/537.36 MMWEBID/3027 MicroMessenger/7.0.18.1740(0x27001235) Process/toolsmp WeChat/arm64 NetType/WIFI Language/zh_CN ABI/arm64",
			"Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; BKK-AL10 Build/HONORBKK-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/10.6 Mobile Safari/537.36",
		},
		"pc": {
			// macOS 10.15.6  Firefox / Chrome / Safari
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:80.0) Gecko/20100101 Firefox/80.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.30 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Safari/605.1.15",
			// Windows 10 Firefox / Chrome / Edge
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.30 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
			// Linux 就算了 ?
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0",
		},
	}
	var realUserAgentList []string
	if ua == "pc" || ua == "mobile" {
		realUserAgentList = userAgentList[ua]
	} else {
		m := userAgentList["mobile"]
		p := userAgentList["pc"]
		realUserAgentList = append(m, p...)
	}

	return realUserAgentList[rand.Intn(len(realUserAgentList))]
}

type URL string

func ReplaceWithReg(reg, src, repl string) string {
	regx := regexp.MustCompile(reg)
	return regx.ReplaceAllString(src, repl)
}

func ReplaceAllWithReg(reg string, src []string, repl string) []string {
	regx := regexp.MustCompile(reg)
	for i, s := range src {
		src[i] = regx.ReplaceAllString(s, repl)
	}
	return src
}

func DoReq(method, URL string, data *map[string]interface{}, options *map[string]interface{}) (*et.Result, error) {
	if options == nil {
		return nil, errors.New("options must be set.")
	}
	if data == nil {
		data = &map[string]interface{}{}
	}
	headers := url.Values{
		"User-Agent": {chooseUserAgent(reflect.ValueOf((*options)["ua"]).String())},
	}
	var form url.Values
	if strings.ToUpper(method) == "POST" {
		headers.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if strings.Contains(string(URL), "music.163.com") {
		headers.Add("Referer", "https://music.163.com")
	}
	if realIP, has := (*options)["realIP"]; has {
		v := reflect.ValueOf(realIP)
		headers.Add("X-Real-IP", v.String())
	}
	var cookieStr string
	var cookies et.Cookies
	if cookie, has := (*options)["cookie"]; has {
		t := reflect.TypeOf(cookie)
		if t.Kind().String() == reflect.Map.String() {
			var cs []string
			v := reflect.ValueOf(cookie)
			keys := v.MapKeys()
			for _, c := range keys {
				cs = append(cs, fmt.Sprintf("%s=%s", url.QueryEscape(c.String()), url.QueryEscape(v.MapIndex(c).String())))
			}
			cookieStr = strings.Join(cs, "; ")
		} else if t.Kind().String() == reflect.TypeOf(et.Cookies{}).Kind().String() {
			cookies = reflect.ValueOf(cookie).Interface().(et.Cookies)
		} else {
			// cookie must be string
			cookieStr = reflect.ValueOf(cookie).String()
		}
	}
	if len(cookieStr) > 0 {
		headers.Add("Cookie", cookieStr)
	}
	hasCookie := func() bool {
		return len(headers.Get("Cookie")) != 0
	}
	if !hasCookie() {
		tv := (*options)["token"]
		if tv != nil {
			headers.Add("Cookie", reflect.ValueOf(tv).String())
		} else {
			headers.Add("Cookie", "")
		}
	}
	if crypto, has := (*options)["crypto"]; has {
		cryptoStr := reflect.ValueOf(crypto).String()
		if cryptoStr == "weapi" {
			if cookieStr == "" {
				cookieStr = cookies.String()
			}
			csrfTokenReg := regexp.MustCompile(`_csrf=([^(;|$)]+)`)
			tokens := csrfTokenReg.FindAllString(cookieStr, -1)
			if len(tokens) > 1 {
				(*data)["csrf_token"] = tokens[1]
			} else {
				(*data)["csrf_token"] = ""
			}
			headers.Set("Cookie", cookieStr)
			form, _ = Weapi(data)
			URL = ReplaceWithReg(`\w*api`, URL, "weapi")
		} else if cryptoStr == "linuxapi" {
			raw := map[string]interface{}{
				"method": method,
				"url":    ReplaceWithReg(`\w*api`, URL, "api"),
				"params": data,
			}
			form, _ = LinuxApi(raw)
			headers.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
			URL = "https://music.163.com/api/linux/forward"
		} else if cryptoStr == "eapi" {
			header := eapiDefaultHeader(&cookies)
			headers.Del("Cookie")
			for k, v := range header {
				headers.Add("Cookie", fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v)))
			}
			(*data)["header"] = header
			form, _ = Eapi(reflect.ValueOf((*options)["url"]).String(), *data)
			URL = ReplaceWithReg(`\w*api`, URL, "eapi")
		}
	}

	if client == nil {
		client = http.DefaultClient
	}
	var resp *http.Response
	var err error
	switch method {
	case http.MethodGet:
		resp, err = GetWithHeaders(URL, form, headers)
		break
		// {"params":"","encSecKey":""}
	case http.MethodPost:
		resp, err = PostFormWithHeaders(URL, form, headers)
		break
	default:
		break
	}
	if resp != nil && err == nil {
		if b, err := ioutil.ReadAll(resp.Body); err == nil {
			regCode := regexp.MustCompile(`\"\w+\"\:400`)
			regMsg := regexp.MustCompile(`\"msg\"\:\"参数错误\"`)
			c := regCode.FindAll(b, -1)
			if len(c) > 0 {
				// fmt.Println(string(c[0]))
				// TODO
			}
			m := regMsg.FindAll(b, -1)
			if len(m) > 0 {
				// fmt.Println(string(m[0]))
				// TODO
			}
			return &et.Result{
				Body:    b,
				Cookies: resp.Cookies(),
				// Cookie: strings.Join(ReplaceAllWithReg(`\s*Domain=[^(;|$)]+;*`, resp.Header.Values("set-cookie"), ""), ";"),
				Header: resp.Header.Clone(),
				Status: resp.StatusCode,
			}, nil
		} else {
			return nil, err
		}
	}
	return nil, errors.New("no data returned")
}

func PostFormWithHeaders(url string, data url.Values, headers url.Values) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		if len(v) > 0 {
			req.Header.Add(k, v[0])
		}
	}
	return client.Do(req)
}

func GetWithHeaders(url string, data url.Values, headers url.Values) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		if len(v) > 0 {
			req.Header.Add(k, v[0])
		}
	}
	return client.Do(req)
}

func eapiDefaultHeader(cookies *et.Cookies) map[string]string {
	dh := map[string]string{}
	if cookies != nil {
		for _, c := range *cookies {
			dh[c.Name] = c.Value
		}
	}
	if _, has := dh[et.CK_OSVER]; !has {
		dh[et.CK_OSVER] = ""
	}
	if _, has := dh[et.CK_DEVICE_ID]; !has {
		dh[et.CK_DEVICE_ID] = ""
	}
	if _, has := dh[et.CK_APPVER]; !has {
		dh[et.CK_APPVER] = "6.1.1"
	}
	if _, has := dh[et.CK_VERSION_CODE]; !has {
		dh[et.CK_VERSION_CODE] = "140"
	}
	if _, has := dh[et.CK_MOBILE_NAME]; !has {
		dh[et.CK_MOBILE_NAME] = ""
	}
	if _, has := dh[et.CK_BUILD_VER]; !has {
		dh[et.CK_BUILD_VER] = string([]rune(time.Now().String())[0:10])
	}
	if _, has := dh[et.CK_RESOLUTION]; !has {
		dh[et.CK_RESOLUTION] = `1920x1080`
	}
	if _, has := dh[et.CK_OS]; !has {
		dh[et.CK_OS] = `android`
	}
	if _, has := dh[et.CK_CHANNEL]; !has {
		dh[et.CK_CHANNEL] = ""
	}
	if _, has := dh[et.CK_REQUEST_ID]; !has {
		dh[et.CK_REQUEST_ID] = fmt.Sprintf("%d", time.Now().UnixNano()/1e2)
	}
	return dh
}
