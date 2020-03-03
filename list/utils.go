package list

import "reflect"

func getTypeName(e interface{}) string {
	return reflect.TypeOf(e).Name()
}
