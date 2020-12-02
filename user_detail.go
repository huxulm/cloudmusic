package cloudmusic

import (
	"fmt"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// UserDetail
// @params q includes of uid string
func UserDetail(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	if res, err := util.DoReq("POST", fmt.Sprintf("%s/%s", USER_DETAIL, q.Get("uid").String()), DefEmptyOpts().Raw(), DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
