package CloudMusic

import (
	"testing"

	"github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	keywords := "摘星"
	expect := 200
	cookies := entities.Cookies{}
	// ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := Search(&cookies, keywords, "", "", ""); err == nil {
		s, err := res.ToSearchRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
}
