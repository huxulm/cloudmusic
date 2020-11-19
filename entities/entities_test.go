package entities

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	qv := QueryValue(reflect.ValueOf(false))
	assert.Equal(t, "false", qv.String())

	qv1 := QueryValue(reflect.ValueOf(64))
	assert.Equal(t, "64", qv1.String())

	qv2 := QueryValue(reflect.ValueOf("queryvalue"))
	assert.Equal(t, "queryvalue", qv2.String())

	qv3 := Query{}
	assert.Equal(t, "", qv3.Get("notexists").String())
}
