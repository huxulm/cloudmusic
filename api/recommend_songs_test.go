package api

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestRecommendSongs(t *testing.T) {
	cookies := ent.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := RecommendSongs(&cookies, nil); err == nil {
		s, err := res.ToRecSongsRes()
		if assert.NoError(t, err) {
			assert.Equal(t, 200, s.Code)
			PersistToFile(s, ".cache/.recommend-songs.yaml")
		}
	}
}
