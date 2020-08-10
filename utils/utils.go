package utils

import "reflect"

func GetTypeName(e interface{}) string {
	if e == nil {
		return NA
	}

	t := reflect.TypeOf(e)

	if res := t.Name(); len(res) != 0 {
		return res
	}

	return t.Elem().Name()
}
