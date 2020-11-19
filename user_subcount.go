package CloudMusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func UserSubcount(cookie string) (*ent.Result, error) {
	if res, err := util.DoReq("POST", USER_SUBCOUNT, DefEmptyOpts().Raw(), DefOpts().CookieString(cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
