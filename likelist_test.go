package CloudMusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestLikeList(t *testing.T) {
	// uid := "<user id>"
	uid := 45964623
	expect := 200
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	query := ent.Query(map[string]interface{}{"uid": uid})
	if res, err := LikeList(&cookies, &query); err == nil {
		s, err := res.ToLikeListRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
			PersistToFile(s, ".likelist.yaml")
		}
	}
}
