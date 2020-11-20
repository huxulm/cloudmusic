package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func SongLyric(cookie *ent.Cookies, id string, br int) (*ent.Result, error) {
	data := map[string]interface{}{
		"id": id,
		"lv": -1,
		"kv": -1,
		"tv": -1,
	}
	if res, err :=
		util.DoReq("POST",
			SONG_LYRIC,
			&data,
			DefOpts().CryptoLinux().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
