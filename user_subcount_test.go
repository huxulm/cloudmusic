package cloudmusic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSubcount(t *testing.T) {
	cookie := "NMTID=00OkR-UFsQ1z-5nh04Mmsoz-yM1080AAAF11LTfXQ; Expires=Fri, 15-Nov-2030 05:37:04 GMT; MUSIC_U=65f5514408fabe018dc222f659bb7fd245690c98fabc9ab19a781434adb85c670931c3a9fbfe3df2; Expires=Wed, 02-Dec-2020 05:37:04 GMT; __csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	res, err := UserSubcount(cookie)
	subcount, err := res.ToSubcountRes()
	assert.Empty(t, err, "convert entity error")
	assert.NotEmpty(t, subcount, "subcount is empty")
	assert.Equal(t, subcount.Code, int64(200), "code should == 200")
}
