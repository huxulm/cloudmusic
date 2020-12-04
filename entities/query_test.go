package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryValue(t *testing.T) {
	q := Query{
		"ids": []string{"OK"},
	}
	assert.Equal(t, "OK", q.Get("ids").Value().([]string)[0])
}
