package utils

import "reflect"

func Struct2Map(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	dataValue := reflect.ValueOf(data)
	for dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		m[dataType.Field(i).Name] = dataValue.Field(i).Interface()
	}
	return m
}
