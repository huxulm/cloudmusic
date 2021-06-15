package api

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// Like adds or deletes a song to or from likelist
// 	q should have a value of key "uid"
func LikeList(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := map[string]interface{}{
		"uid": q.Get("uid").String(),
	}
	if res, err :=
		util.DoReq("POST", LIKE_LIST, &data, DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
