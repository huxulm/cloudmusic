package CloudMusic

import (
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/jackdon/cloudmusic/util"
)

func RecommendResource(cookie *ent.Cookies, q *ent.Query) (*ent.Result, error) {
	if res, err :=
		util.DoReq("POST", RECOMMEND_RESOURCE, DefEmptyOpts().Raw(), DefOpts().Cookie(*cookie).Raw()); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
