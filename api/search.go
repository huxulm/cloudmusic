package api

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

// Search
// @params q includes of keywords, type_, limit, offset string
func Search(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	// type: // 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频
	data := map[string]interface{}{
		"s":      q.Get("keywords").StrDefault(""),
		"type":   q.Get("type").StrDefault("1"),
		"limit":  q.Get("limit").StrDefault("30"),
		"offset": q.Get("offset").StrDefault("0"),
	}
	if res, err :=
		util.DoReq("POST",
			SEARCH,
			&data,
			DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
