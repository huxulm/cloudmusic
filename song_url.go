package cloudmusic

import (
	"fmt"
	"strings"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// SongURL
// @params q includes of ids []string, br int
func SongURL(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := map[string]interface{}{
		"br": q.Get("br").IntDefault(999000),
	}
	ids := q.Get("ids").Value().([]string)
	data["ids"] = fmt.Sprintf("[%s]", strings.Join(ids, ","))
	url := "/api/song/enhance/player/url"
	if res, err :=
		util.DoReq("POST",
			SONG_URL,
			&data,
			DefOpts().CryptoEapi().Cookie(*cookie).URL(url).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
