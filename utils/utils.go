package utils

import "reflect"

func GetTypeName(e interface{}) string {
	if e == nil {
		return NA
	}
	return reflect.TypeOf(e).Name()
}
