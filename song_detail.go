package cloudmusic

import (
	"errors"
	"fmt"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// SongDetail
func SongDetail(cookie *ent.Cookies, q *ent.Query) (res *ent.Result, err error) {
	data := map[string]interface{}{
		"c":   fmt.Sprintf("[%s]", q.Get("c").String()),
		"ids": fmt.Sprintf("[%s]", q.Get("ids").String()),
	}
	defer func() {
		if r := recover(); r != nil {
			res, err = nil, errors.New(fmt.Sprintf("%v", r))
		}
	}()
	if res, err :=
		util.DoReq("POST",
			SONG_DETAIL,
			&data,
			DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
