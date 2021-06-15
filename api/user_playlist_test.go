package api

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserPlaylist(t *testing.T) {
	uid := "45964623"
	var cookies ent.Cookies
	ParseFromFile(&cookies, ".cookies.yaml")
	res, err := UserPlaylist(&cookies, &ent.Query{"uid": uid, "limit": nil, "offset": nil})
	assert.Empty(t, err, "convert entity error")
	pl, err := res.ToPlaylistRes()
	assert.NotEmpty(t, pl, "playlist is empty")
	assert.Equal(t, pl.Code, 200, "code should == 200")
	PersistToFile(pl, ".playlist.yaml")
}
