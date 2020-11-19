package CloudMusic

import (
	"fmt"
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestPlaylistDetai(t *testing.T) {
	pid := "40698703"
	// cookie := "__csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	var cookies ent.Cookies
	ParseFromFile(&cookies, ".cookies.yaml")
	assert.Empty(t, cookies)
	expect := 6
	if res, err := PlaylistDetail(&cookies, pid, nil); err == nil {
		fmt.Println(res.BodyAsString())
		d, err := res.ToPlaylistDetailRes()
		assert.Error(t, err)
		assert.Equal(t, expect, len(d.Playlist.Tracks))
		PersistToFile(d, ".playlist-detail.yaml")
	}
}
