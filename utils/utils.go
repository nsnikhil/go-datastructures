package utils

import "reflect"

func GetTypeName(e interface{}) string {
	if e == nil {
		return na
	}
	return reflect.TypeOf(e).Name()
}
