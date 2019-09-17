package format

import (
	"reflect"
	"encoding/json"
)

func ClearStruct(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}


func StructToStr(content interface{}) string {
	str, _ := json.Marshal(content)
	return string(str)
}
