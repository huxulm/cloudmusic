package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func LoginCellphone(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := &map[string]interface{}{
		"phone":         q.Get("phone").String(),
		"countrycode":   "86",
		"password":      util.Md5String(q.Get("password").String()),
		"rememberLogin": "true",
	}
	if res, err := util.DoReq("POST", LOGIN_CELLPHONE, data, DefOpts().Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
