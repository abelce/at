package structUtils

import "reflect"

// data必须是struct，map，slice， d必须是指针的形式
func SetStructValueFromMap(data interface{}, inStructPtr interface{}) {
	if reflect.TypeOf(inStructPtr).Kind() == reflect.Ptr {
		rType := reflect.TypeOf(inStructPtr).Elem()
		rVal := reflect.ValueOf(inStructPtr).Elem()

		dataValue := reflect.ValueOf(data)

		for i := 0; i < rType.NumField(); i++ {
			t := rType.Field(i)
			f := rVal.Field(i)
			// 通过filedname获取值
			fileValue := dataValue.FieldByName(t.Name)
			if fileValue.IsValid() {
				f.Set(fileValue)
			}
		}
	}
}
