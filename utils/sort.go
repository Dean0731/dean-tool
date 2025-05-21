package utils

import (
	"fmt"
	"reflect"
	"sort"
)

func SortByField[T any](data []T, fieldName string) {
	v := reflect.ValueOf(data)
	if v.Len() == 0 {
		return
	}

	elemType := v.Type().Elem()
	field, ok := elemType.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("找不到名为 %s 的字段\n", fieldName))
		return
	}
	fieldIndex := field.Index[0]

	sort.SliceStable(data, func(i, j int) bool {
		x := v.Index(i).Field(fieldIndex)
		y := v.Index(j).Field(fieldIndex)
		switch x.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return x.Int() < y.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return x.Uint() < y.Uint()
		case reflect.Float32, reflect.Float64:
			return x.Float() < y.Float()
		case reflect.String:
			return x.String() < y.String()
		default:
			panic("不支持的字段类型")
		}
	})
}
