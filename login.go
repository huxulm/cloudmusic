package CloudMusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func LoginCellphone(phone, password string) (*ent.Result, error) {
	if res, err := util.DoReq("POST", LOGIN_CELLPHONE, &map[string]interface{}{
		"phone":         phone,
		"countrycode":   "86",
		"password":      util.Md5String(password),
		"rememberLogin": "true",
	}, DefOpts().Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
