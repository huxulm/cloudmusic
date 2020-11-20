package cloudmusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func PlaylistHqTags(cookie *ent.Cookies) (*ent.Result, error) {
	if res, err :=
		util.DoReq("POST",
			PLAYLIST_HIGHQUALITY_TAGS,
			DefEmptyOpts().Raw(),
			DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
