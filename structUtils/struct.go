package structUtils

// data是map， d必须是指针的形式
func SetStructValueFromMap( data map[string]interace{}, inStructPtr interface{}) {
    if reflect.TypeOf(inStructPtr).Kind() == reflect.Ptr {
		rType := reflect.TypeOf(inStructPtr).Elem()
		rVal := reflect.ValueOf(inStructPtr).Elem()
	
		for i := 0; i < rType.NumField(); i++ {
			t := rType.Field(i)
			f := rVal.Field(i)
	
			if v, ok := data[t.Name]; ok {
				f.Set(reflect.ValueOf(v))
			}
		}
	}
}