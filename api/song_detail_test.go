package api

import (
	"fmt"
	"strings"
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestSongDetail(t *testing.T) {
	ids := []string{"40021474", "159107", "34017157"}
	query := buildSongDetailQuery(ids)
	expect := 3
	cookies := ent.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	res, err := SongDetail(&cookies, query)
	if assert.NoError(t, err) {
		s, err := res.ToSongDetailRes()
		if assert.NoError(t, err) {
			assert.Equal(t, expect, len(s.Songs))
		}
	}
}

func buildSongDetailQuery(ids []string) *ent.Query {
	query := ent.Query(map[string]interface{}{"ids": strings.Join(ids, ",")})
	c := make([]string, len(ids))
	for i, id := range ids {
		c[i] = fmt.Sprintf(`{"id":%s}`, id)
	}
	query["c"] = strings.Join(c, ",")
	return &query
}
