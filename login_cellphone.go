package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func Login(email, password string) (*ent.Result, error) {
	if res, err := util.DoReq("POST", LOGIN_EMAIL, &map[string]interface{}{
		"username":      email,
		"password":      util.Md5String(password),
		"rememberLogin": "true",
	}, DefOpts().Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
