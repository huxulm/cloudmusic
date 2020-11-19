package CloudMusic

import (
	"fmt"
	"strings"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func SongURL(cookie *ent.Cookies, ids []string, br int) (*ent.Result, error) {
	data := map[string]interface{}{}
	data["ids"] = fmt.Sprintf("[%s]", strings.Join(ids, ","))
	if br > 0 {
		data["br"] = br
	} else {
		data["br"] = 999000
	}
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
