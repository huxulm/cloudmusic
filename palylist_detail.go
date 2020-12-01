package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func PlaylistDetail(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := map[string]interface{}{
		"id": q.Get("id").String(),
		"n":  100000,
	}

	if s := q.Get("s"); s == nil {
		data["s"] = 8
	} else {
		data["s"] = *s.Int64()
	}
	if res, err := util.DoReq("POST", PLAYLIST_DETAIL, &data, DefOpts().CryptoLinux().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
