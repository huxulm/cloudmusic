package util

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestRequest(t *testing.T) {
	phone := os.Getenv("GCM_PHONE")
	pass := os.Getenv("GCM_PASS")
	if phone == "" {
		phone = "18616777015"
	}
	if pass == "" {
		phone = "<your pass>"
	}
	data := map[string]interface{}{
		"phone":         phone,
		"countrycode":   "86",
		"password":      Md5String(pass),
		"rememberLogin": "true",
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"ua":     "pc",
		"cookie": map[string]string{"os": "pc"},
		// "realIP": "<IP>",
	}
	if result, err := DoReq("POST", `https://music.163.com/weapi/login/cellphone`, &data, &options); err == nil {
		fmt.Println("Body:")
		fmt.Println(result.BodyAsString())
		fmt.Printf("Cookie:\n%s", result.Cookies.String())
		h := map[string][]string(result.Header)
		if hb, err := json.Marshal(h); err == nil {
			fmt.Printf("Headers:\n%s", hb)
		}
	}
}
