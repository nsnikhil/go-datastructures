package list

import "reflect"

func getTypeName(e interface{}) string {
	if e == nil {
		return na
	}
	return reflect.TypeOf(e).Name()
}
