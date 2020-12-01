package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func Login(cookies *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := &map[string]interface{}{
		"username":      q.Get("email").String(),
		"password":      util.Md5String(q.Get("password").String()),
		"rememberLogin": "true",
	}
	if res, err := util.DoReq("POST", LOGIN_EMAIL, data, DefOpts().Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
