package main

import (
	"fmt"
	"reflect"
	"testing"

	jsonHelper "github.com/OskarEek/llmgo/utilities/jsonHelper"
)

type PersonTest struct {
	Name     string
	Age      int
	Location string
}

func TestStructJson(t *testing.T) {
	var obj PersonTest
	structure, err := jsonHelper.GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}

func TestStructSliceJson(t *testing.T) {
	var obj []PersonTest
	structure, err := jsonHelper.GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}

func TestSliceJson(t *testing.T) {
	var obj []int
	structure, err := jsonHelper.GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}
