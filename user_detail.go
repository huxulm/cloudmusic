package cloudmusic

import (
	"fmt"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func UserDetail(uid, cookie string) (*ent.Result, error) {
	if res, err := util.DoReq("POST", fmt.Sprintf("%s/%s", USER_DETAIL, uid), DefEmptyOpts().Raw(), DefOpts().CookieString(cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
