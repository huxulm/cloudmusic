package cloudmusic

import (
	"fmt"
	"os"
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestLoginCellphone(t *testing.T) {
	phone := os.Getenv("GCM_PHONE")
	if phone == "" {
		phone = "<your phone>"
	}
	pass := os.Getenv("GCM_PASS")
	if pass == "" {
		pass = "<your password>"
	}
	nickname := os.Getenv("GCM_NICK_NAME")
	query := ent.Query(map[string]interface{}{"phone": phone, "password": pass})
	if res, err := LoginCellphone(nil, &query); err == nil {
		fmt.Println(res.Cookies.String())
		if lcr, err := res.ToLoginCellRes(); err != nil {
			assert.Fail(t, "failed:", res.AsJSON(), err.Error())
		} else {
			fmt.Printf("logged: %s", lcr.Profile.Nickname)
			assert.Equal(t, lcr.Profile.Nickname, nickname, "nickname not correct")
		}
	} else {
		assert.Fail(t, "login_cell test failed", err)
	}
}
