package utils

import (
	"encoding/json"
	"fmt"
	"github.com/kube-stack/multicloud_service/src/utils"
	"reflect"
	"testing"
)

type Animal struct {
	Name string
	Kind string
}

func TestConstructStruct(t *testing.T) {
	a := Animal{
		Name: "animalA",
		Kind: "KindA",
	}
	jsonBytes, _ := json.Marshal(a)
	b := &Animal{}
	err := utils.ConstructStructOfPtr(b, jsonBytes)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(a, *b) {
		t.Error("a and b are not equal")
	}
	c := reflect.New(reflect.TypeOf(a))
	err = utils.ConstructStructOfValue(c, jsonBytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.Interface())
	//notice c is a pointer, so a need with  &
	if !reflect.DeepEqual(&a, c.Interface()) {
		t.Error("a and c are not equal")
	}
}
