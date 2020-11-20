package cloudmusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestLike(t *testing.T) {
	id := "<song id>"
	like := false
	expect := 200
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	query := ent.Query(map[string]interface{}{"id": id, "like": like})
	if res, err := Like(&cookies, &query); err == nil {
		s, err := res.ToLikeRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
}
