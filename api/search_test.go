package api

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	keywords := "周杰伦"
	type_ := "100"
	expect := 200
	cookies := ent.Cookies{}
	// search singers
	if res, err := Search(&cookies, &ent.Query{"keywords": keywords, "type": type_}); err == nil {
		s, err := res.ToSearchSingerRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
	// search songs
	type_ = "1"
	if res, err := Search(&cookies, &ent.Query{"keywords": keywords, "type": type_}); err == nil {
		s, err := res.ToSearchRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, s.Code)
		}
	}
}
