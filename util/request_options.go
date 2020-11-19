package util

import "sync"

type RequestOptions map[string]interface{}

var lock sync.Mutex

// safe to use in multiple goroutines
func (opt *RequestOptions) SetOption(key string, value interface{}) {
	lock.Lock()
	(*opt)[key] = value
	lock.Unlock()
}

func (opt *RequestOptions) UA(ua string) *RequestOptions {
	opt.SetOption("ua", ua)
	return opt
}

func (opt *RequestOptions) RealIP(ip string) *RequestOptions {
	opt.SetOption("realIP", ip)
	return opt
}

func (opt *RequestOptions) Cookie(cookie interface{}) *RequestOptions {
	// cookie can be set with a map
	opt.SetOption("cookie", cookie)
	return opt
}

func (opt *RequestOptions) CookieString(cookie string) *RequestOptions {
	// cookie can be set with a string
	opt.SetOption("cookie", cookie)
	return opt
}

func (opt *RequestOptions) URL(url string) *RequestOptions {
	// cookie can be set with a string
	opt.SetOption("url", url)
	return opt
}

func (opt *RequestOptions) CryptoLinux() *RequestOptions {
	opt.SetOption("crypto", "linuxapi")
	return opt
}
func (opt *RequestOptions) CryptoEapi() *RequestOptions {
	opt.SetOption("crypto", "eapi")
	return opt
}

func (opt *RequestOptions) Raw() *map[string]interface{} {
	raw := map[string]interface{}(*opt)
	return &raw
}

func DefaultRequestOpts() *RequestOptions {
	return &RequestOptions{
		"crypto": "weapi",
		"ua":     "pc",
		"cookie": map[string]string{"os": "pc"},
		// "realIP": "47.98.170.133",
	}
}

func DefaultEmptyRequestOpts() *RequestOptions {
	return &RequestOptions{}
}
