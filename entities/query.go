package entities

import (
	"fmt"
	"reflect"
)

type Query map[string]interface{}
type QueryValue reflect.Value

func (qv *QueryValue) String() string {
	if qv == nil {
		return ""
	}
	v := reflect.Value(*qv)
	return fmt.Sprintf("%v", v.Interface())
}
func (qv *QueryValue) Bool() *bool {
	if qv == nil {
		return nil
	}
	v := reflect.Value(*qv).Bool()
	return &v
}
func (qv *QueryValue) Int64() *int64 {
	if qv == nil {
		return nil
	}
	v := reflect.Value(*qv).Int()
	return &v
}
func (qv *QueryValue) Value() interface{} {
	if qv == nil {
		return nil
	}
	v := reflect.Value(*qv).Interface()
	return &v
}

func (qv *QueryValue) StrDefault(defa string) string {
	if qv == nil {
		if len(defa) == 0 {
			return ""
		} else {
			return defa
		}
	}
	v := reflect.Value(*qv)
	return fmt.Sprintf("%v", v.Interface())
}

func (q *Query) Get(key string) *QueryValue {
	if v, has := (*q)[key]; has {
		qv := QueryValue(reflect.ValueOf(v))
		return &qv
	}
	return nil
}
