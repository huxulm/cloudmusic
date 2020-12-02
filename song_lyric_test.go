package cloudmusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestSongLyric(t *testing.T) {
	id := "1447981499"
	expect := 200
	cookies := entities.Cookies{}
	query := entities.Query{"id": id}
	// ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := SongLyric(&cookies, &query); err == nil {
		s, err := res.ToSongLyricRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
}
