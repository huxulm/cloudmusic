package cloudmusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserDetai(t *testing.T) {
	uid := "45964623"
	// cookie := "__csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	expect := "<your nickname>"
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := UserDetail(uid, cookies.FindByKey("__csrf").String()); err == nil {
		udr, err := res.ToUserDetailRes()
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, expect, udr.Profile.Nickname)
	}
}
