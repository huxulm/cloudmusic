package CloudMusic

import "github.com/jackdon/cloudmusic/util"

type D map[string]interface{}

func DefOpts() *util.RequestOptions {
	return util.DefaultRequestOpts()
}

func DefEmptyOpts() *util.RequestOptions {
	return util.DefaultEmptyRequestOpts()
}

func Logout(cookie string) error {
	util.DoReq("POST", LOGOUT, nil, DefOpts().CookieString(cookie).Raw())
	return nil
}
