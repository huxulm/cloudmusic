package cloudmusic

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserPlaylist(t *testing.T) {
	uid := "45964623"
	// cookie := "NMTID=00OkR-UFsQ1z-5nh04Mmsoz-yM1080AAAF11LTfXQ; Expires=Fri, 15-Nov-2030 05:37:04 GMT; MUSIC_U=65f5514408fabe018dc222f659bb7fd245690c98fabc9ab19a781434adb85c670931c3a9fbfe3df2; Expires=Wed, 02-Dec-2020 05:37:04 GMT; __csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	var cookies ent.Cookies
	ParseFromFile(&cookies, ".cookies.yaml")
	res, err := UserPlaylist(&cookies, uid, nil, nil)
	assert.Empty(t, err, "convert entity error")
	pl, err := res.ToPlaylistRes()
	assert.NotEmpty(t, pl, "playlist is empty")
	assert.Equal(t, pl.Code, 200, "code should == 200")
	PersistToFile(pl, ".playlist.yaml")
}
