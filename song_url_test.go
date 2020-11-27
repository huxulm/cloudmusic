package cloudmusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestSongURL(t *testing.T) {
	ids := []string{"40021474,421934070,34017157"}
	expect := 3
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := SongURL(&cookies, ids, -1); err == nil {
		s, err := res.ToSongURLRes()
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, expect, len(s.Data))
	}
}
