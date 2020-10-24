package utils

import (
	"github.com/nsnikhil/go-datastructures/liberr"
	"reflect"
)

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

//TODO: RENAME
func AreAllSameType(data ...interface{}) error {
	typeURL := GetTypeName(data[0])
	sz := len(data)

	for i := 1; i < sz; i++ {
		if et := GetTypeName(data[i]); et != typeURL {
			return liberr.TypeMismatchError(typeURL, et)
		}
	}

	return nil
}
