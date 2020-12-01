package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

type D map[string]interface{}

func DefOpts() *util.RequestOptions {
	return util.DefaultRequestOpts()
}

func DefEmptyOpts() *util.RequestOptions {
	return util.DefaultEmptyRequestOpts()
}

func Logout(cookie *ent.Cookies, q *ent.Query) error {
	util.DoReq("POST", LOGOUT, nil, DefOpts().Cookie(*cookie).Raw())
	return nil
}
