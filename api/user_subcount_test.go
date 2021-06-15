package api

import (
	"testing"

	ent "github.com/jackdon/cloudmusic/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserSubcount(t *testing.T) {
	var cookies ent.Cookies
	ParseFromFile(&cookies, ".cookies.yaml")
	res, err := UserSubcount(&cookies, nil)
	subcount, err := res.ToSubcountRes()
	assert.Empty(t, err, "convert entity error")
	assert.NotEmpty(t, subcount, "subcount is empty")
	assert.Equal(t, subcount.Code, int64(200), "code should == 200")
}
