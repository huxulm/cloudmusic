package cloudmusic

import (
	"fmt"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// Like adds or deletes a song to or from likelist
// the string id is an id of song
func Like(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	data := map[string]interface{}{
		"trackId": q.Get("id").String(),
		"like":    q.Get("like").Bool(),
	}
	if res, err :=
		util.DoReq("POST",
			fmt.Sprintf("%s?alg=%s&trackId=%s&time=%s", LIKE, q.Get("alg").StrDefault("itembased"), q.Get("id").String(), q.Get("time").StrDefault("25")),
			&data,
			DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
