package utils

import (
	"fmt"
	"reflect"
)

/*
this function is used to call create request function in functionRegistry
*/
func CallFunction(funcName string, functionRegistry map[string]interface{}, params ...interface{}) (result []interface{}, err error) {
	funcInterface, ok := functionRegistry[funcName]
	if !ok {
		err = fmt.Errorf("can't find function %s", funcName)
		return
	}
	f := reflect.ValueOf(funcInterface)
	if len(params) != f.Type().NumIn() {
		err = fmt.Errorf("The number of params is out of index.")
		return
	}
	//construct prarmeter
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	//construct return value
	numOut := f.Type().NumOut()
	var res []reflect.Value
	result = make([]interface{}, f.Type().NumOut())
	res = f.Call(in)
	for i := 0; i < numOut; i++ {
		result[i] = res[i].Interface()
	}
	return
}

/*
this function is used to call a struct's method
*/
func CallMethod(instance interface{}, methodName string, params ...interface{}) (result []interface{}, err error) {

	method := reflect.ValueOf(instance).MethodByName(methodName)

	if !method.IsValid() {
		err = fmt.Errorf("instance %v doesn't have method %s\n", reflect.ValueOf(instance), methodName)
	}
	in := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		in[i] = reflect.ValueOf(params[i])
	}
	ret := method.Call(in)
	for i := 0; i < method.Type().NumOut(); i++ {
		result = append(result, ret[i].Interface())
	}
	return
}
