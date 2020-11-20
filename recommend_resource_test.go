package cloudmusic

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestRecommendResouce(t *testing.T) {
	cookies := ent.Cookies{}
	ParseFromFile(&cookies, ".cookies.yaml")
	if res, err := RecommendResource(&cookies, nil); err == nil {
		s, err := res.ToRecResourceRes()
		if assert.NoError(t, err) {
			assert.Equal(t, 200, s.Code)
			PersistToFile(s, ".cache/.recommend-resource.yaml")
		}
	}
}
