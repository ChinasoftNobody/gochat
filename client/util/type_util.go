package util

import (
	"reflect"
)

/**
判断类型是否为目标类型
*/
func IsTypeOf(args interface{}, t interface{}) bool {
	return reflect.TypeOf(args) == reflect.TypeOf(t)
}
