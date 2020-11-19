package CloudMusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func Search(cookie *ent.Cookies, keywords, type_, limit, offset string) (*ent.Result, error) {
	// type: // 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频
	if len(type_) == 0 {
		type_ = "1"
	}
	if len(limit) == 0 {
		limit = "30"
	}
	if len(offset) == 0 {
		offset = "0"
	}
	data := map[string]interface{}{
		"s":      keywords,
		"type":   type_,
		"limit":  limit,
		"offset": offset,
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
