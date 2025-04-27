package jsonhelper

import (
	"reflect"
	"testing"
)

type ParentTest struct {
	ParentName string
	ParentAge  int
	ParentTo   string
}

type PersonTest struct {
	Name     string
	Age      int
	Location string
	Parent   *ParentTest
}

type FaultyStruct struct {
	selfReference *FaultyStruct
}

func TestFaultyStruct(t *testing.T) {
	var obj FaultyStruct
	structure, err := GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		t.Logf("%s", err)
	} else {
		t.Logf("%s", structure)
		t.FailNow()
	}
}

func TestStructJson(t *testing.T) {
	var obj PersonTest
	structure, err := GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}

func TestStructSliceJson(t *testing.T) {
	var obj []PersonTest
	structure, err := GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}

func TestSliceJson(t *testing.T) {
	var obj []int
	structure, err := GetJsonStructureFromType(reflect.TypeOf(obj))
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	t.Logf("\n%s", structure)
}
