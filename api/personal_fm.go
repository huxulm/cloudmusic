package api

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func PersonalFM(cookie *ent.Cookies, query *ent.Query) (*ent.Result, error) {
	if res, err := util.DoReq("POST", PERSONAL_FM, DefEmptyOpts().Raw(),
		DefOpts().CryptoLinux().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
