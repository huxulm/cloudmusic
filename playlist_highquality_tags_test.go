package cloudmusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestPlaylistHqTags(t *testing.T) {
	expect := 200
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := PlaylistHqTags(&cookies); err == nil {
		s, err := res.ToHqTagsRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
}
