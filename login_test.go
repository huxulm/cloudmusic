package cloudmusic

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLogin(t *testing.T) {
	email := os.Getenv("GCM_EMAIL")
	if email == "" {
		email = "<your email>"
	}
	pass := os.Getenv("GCM_PASS")
	if pass == "" {
		pass = "<your pass>"
	}
	nickname := os.Getenv("GCM_NICK_NAME")
	if nickname == "" {
		nickname = "<your nickname>"
	}
	query := ent.Query(map[string]interface{}{"email": email, "password": pass})
	if res, err := Login(nil, &query); err == nil {
		fmt.Println(res.Cookies.String())
		if lcr, err := res.ToLoginCellRes(); err != nil {
			assert.Fail(t, "failed:", res.AsJSON(), err.Error(), res.BodyAsString())
		} else {
			PersistToFile(lcr, ".login.yaml")
			PersistToFile(map[string][]string(res.Header), ".headers.yaml")
			PersistToFile(res.Cookies, ".cookies.yaml")
			assert.Equal(t, lcr.Profile.Nickname, nickname, "nickname not correct")
		}
	} else {
		assert.Fail(t, "login_cell test failed", err)
	}
}

func PersistToFile(v interface{}, file string) {
	if out, err := yaml.Marshal(v); err == nil {
		ioutil.WriteFile(filepath.Join("./", file), out, os.ModePerm)
	}
}

func ParseFromFile(v interface{}, file string) error {
	if b, err := ioutil.ReadFile(filepath.Join("./", file)); err == nil {
		return yaml.Unmarshal(b, v)
	} else {
		return err
	}
}
