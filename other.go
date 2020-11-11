package dzgutils

/*
 other utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"fmt"
)

/*
  ternary operator, replace other language: a == b ? c : d
*/
func IIIInterfaceOperator(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func InterfaceToString(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	}
	return ""
}

func InterfaceToInt(i interface{}) int {
	switch t := i.(type) {
	case int:
		return t
	}
	return 0
}

// StructToMap : struct转为map
func StructToMap(source interface{}) map[string]interface{}{
	// 判断，interface转为interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Struct {
		panic("ERROR: Unknown type, struct expected.")
	}

	elem := v.Interface()

	data := make(map[string]interface{})
	objT := reflect.TypeOf(elem)
	objV := reflect.ValueOf(elem)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	return data
}

// StructSliceToMapSlice : struct切片转为map切片
func StructSliceToMapSlice(source interface{}) []map[string]interface{}{
	// 判断，interface转为[]interface{}
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	// 转换之后的结果变量
	res := make([]map[string]interface{}, 0)

	// 通过遍历，每次迭代将struct转为map
	for _, elem := range ret {
		data := make(map[string]interface{})
		objT := reflect.TypeOf(elem)
		objV := reflect.ValueOf(elem)
		for i := 0; i < objT.NumField(); i++ {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
		res = append(res, data)
	}
	return res
}

func FmtPrintln(info interface{}) {
	if info != nil {
		switch t := info.(type) {
		case struct{}:
			fmt.Println(StructToJson(t))
		default:
			fmt.Println(t)
		}
	} else {
		fmt.Println(nil)
	}
}
