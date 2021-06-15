package api

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// UserPlaylist q includes of uid string, limit, offset int
func UserPlaylist(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := map[string]interface{}{
		"uid":    q.Get("uid").String(),
		"limit":  q.Get("limit").IntDefault(30),
		"offset": q.Get("offset").IntDefault(0),
	}
	data["includeVideo"] = true
	if res, err := util.DoReq("POST", USER_PLAYLIST, &data, DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
