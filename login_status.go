package CloudMusic

import (
	"regexp"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func LoginStatus(cookie string) (*ent.Result, error) {
	if res, err := util.DoReq("GET", LOGIN_STATUS,
		&map[string]interface{}{},
		DefEmptyOpts().CookieString(cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}

var ProfileReg = regexp.MustCompile(`GUser\s*=\s*([^;]+);`)
var BindingsReg = regexp.MustCompile(`GBinds\s*=\s*([^;]+);`)

func findProfileContent(body []byte) []byte {
	if all := ProfileReg.FindAll(body, -1); len(all) > 0 {
		return all[0]
	}
	return nil
}
func findBindingsContent(body []byte) []byte {
	if all := BindingsReg.FindAll(body, -1); len(all) > 0 {
		return all[0]
	}
	return nil
}
