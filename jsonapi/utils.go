package jsonapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"

	"errors"

	"strings"
)

var (
	//ErrCannotUnmarshalCommand @todo 没有细化错误
	ErrCannotUnmarshalCommand = errors.New("UnmarshalCommand error")
)

//UnmarshalCommand 帮助反序列化命令，因为协商at所使用的jsonapi中的command类型仅为基础类型，故而将relationships拆分为xxType xxID形式
func UnmarshalCommand(in io.Reader, v interface{}) (e error) {
	data, e := ioutil.ReadAll(in)
	if e != nil {
		return e
	}

	var payload map[string]interface{}
	e = json.Unmarshal(data, &payload)
	if e != nil {
		return e
	}

	pdata, ok := payload["data"].(map[string]interface{})
	//非data结构的采用原始数据去进行命令的初始化
	if !ok {
		e = json.Unmarshal(data, v)
		if e == nil {
			return nil
		}

		return e
	}
	if pdata["attributes"] != nil {
		attributes, ok := pdata["attributes"].(map[string]interface{})
		if !ok {
			return ErrCannotUnmarshalCommand
		}
		data, e = json.Marshal(attributes)
		if e != nil {
			return e
		}
		e = json.Unmarshal(data, v)
		if e != nil {
			return e
		}
	}
	if pdata["relationships"] != nil {
		relationships, ok := pdata["relationships"].(map[string]interface{})
		if !ok {
			return ErrCannotUnmarshalCommand
		}

		//动态赋值
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Ptr && !value.Elem().CanSet() {
			return ErrCannotUnmarshalCommand
		}
		value = value.Elem()

		for key, relation := range relationships {
			key = strings.ToUpper(key[0:1]) + key[1:]
			typeName := key + "Type"
			idName := key + "ID"
			ptype := value.FieldByName(typeName)
			pid := value.FieldByName(idName)
			if ptype.IsValid() && ptype.CanSet() {
				rType := relation.(map[string]interface{})["data"].(map[string]interface{})["type"].(string)
				ptype.SetString(rType)
			}
			if pid.IsValid() && pid.CanSet() {
				rID := relation.(map[string]interface{})["data"].(map[string]interface{})["id"].(string)
				pid.SetString(rID)
			}
		}
	}
	return nil
}
