package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func UserPlaylist(cookie *ent.Cookies, uid string, limit, offset *int) (*ent.Result, error) {
	data := map[string]interface{}{}
	data["uid"] = uid
	if limit == nil {
		data["limit"] = 30
	} else {
		data["limit"] = *limit
	}
	if offset == nil {
		data["offset"] = 0
	} else {
		data["offset"] = *offset
	}
	data["includeVideo"] = true
	if res, err := util.DoReq("POST", USER_PLAYLIST, &data, DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
