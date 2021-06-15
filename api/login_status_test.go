package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginStatus(t *testing.T) {
	cookie := "NMTID=00OkR-UFsQ1z-5nh04Mmsoz-yM1080AAAF11LTfXQ; Expires=Fri, 15-Nov-2030 05:37:04 GMT; Path=/;MUSIC_U=65f5514408fabe018dc222f659bb7fd245690c98fabc9ab19a781434adb85c670931c3a9fbfe3df2; Expires=Wed, 02-Dec-2020 05:37:04 GMT; Path=/;__remember_me=true; Expires=Wed, 02-Dec-2020 05:37:04 GMT; Path=/;__csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	if r, err := LoginStatus(cookie); err == nil {
		fmt.Println(string(findProfileContent(r.Body)))
		fmt.Println(r.Status)
	} else {
		assert.Fail(t, err.Error())
	}
}

func TestJSONParse(t *testing.T) {
	_ = fmt.Sprintf("eval('%s')", `{userId:45964623,nickname:"<your nickname>",avatarUrl:"http://p4.music.126.net/uQv7-TsW88LaQYBwmkN5Kg==/2907108744406549.jpg\",birthday:"694195200000",userType:0,djStatus:0}`)
}
