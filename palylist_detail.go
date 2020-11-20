package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func PlaylistDetail(cookie *ent.Cookies, id string, s *int) (*ent.Result, error) {
	data := map[string]interface{}{}
	data["id"] = id
	data["n"] = 100000
	if s == nil {
		data["s"] = 8
	} else {
		data["s"] = *s
	}
	if res, err := util.DoReq("POST", PLAYLIST_DETAIL, &data, DefOpts().CryptoLinux().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
