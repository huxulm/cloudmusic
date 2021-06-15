package api

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestPersonalFM(t *testing.T) {
	expect := 200
	cookies := entities.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := PersonalFM(&cookies, nil); err == nil {
		s, err := res.ToPersonalFMRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
			PersistToFile(s, ".likelist.yaml")
		}
	}
}
