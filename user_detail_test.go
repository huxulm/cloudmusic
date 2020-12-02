package cloudmusic

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserDetai(t *testing.T) {
	uid := "45964623"
	// cookie := "__csrf=0a85b7c137e91ab7429a0525120878be; Expires=Wed, 02-Dec-2020 05:37:14 GMT; Path=/"
	expect := "<your nickname>"
	cookies := ent.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := UserDetail(&cookies, &ent.Query{"uid": uid}); err == nil {
		udr, err := res.ToUserDetailRes()
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, expect, udr.Profile.Nickname)
	}
}
