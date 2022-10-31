package utils

import (
	"cloud_manager/src/utils"
	"fmt"
	"reflect"
	"testing"
)

type RetStruct struct {
	Name  string
	Kind  string
	Value string
}

func (a *Animal) GetStruct(value string) *RetStruct {
	return &RetStruct{
		Name:  a.Name,
		Kind:  a.Kind,
		Value: value,
	}
}
func TestCallMethod(t *testing.T) {
	a := &Animal{
		Name: "Animala",
		Kind: "Kinda",
	}
	result, err := utils.CallMethod(a, "GetStruct", "valuea")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(reflect.ValueOf(result[0]).Type())        //*utils.ResStruct
	fmt.Println(reflect.ValueOf(result[0]).Type().Kind()) //ptr
	fmt.Printf("%v", reflect.ValueOf(result[0]))          //working

}

func TestMap(t *testing.T) {
	b := map[string]string{
		"a": "a",
	}
	fmt.Println(b["a"])
	c := b["b"]
	fmt.Printf(c)
}
