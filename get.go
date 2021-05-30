package at

import (
	"fmt"
	"reflect"
	"strings"

	"abelce/at/errors"
)

func GetWapper(st interface{}) func(path string) (interface{}, error) {
	return func(path string) (interface{}, error) {
		keySlice := strings.Split(path, ".")
		v := reflect.ValueOf(st)

		for _, key := range keySlice {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			if v.Kind() != reflect.Struct {
				return nil, errors.New(fmt.Sprintf("only accept structs: %T", v))
			}

			v = v.FieldByName(key)
		}

		return v, nil
	}
}
