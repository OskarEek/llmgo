package jsonhelper

import (
	"fmt"
	"reflect"
)

func GetJsonStructureFromType(mapType reflect.Type) (string, error) {
	var jsonFormat = ""

	switch mapType.Kind() {
	case reflect.Struct:
		structure, err := GetJsonStructureFromStruct(mapType)
		if err != nil {
			return "", err
		}
		jsonFormat += structure
	case reflect.Slice:
		structure, err := GetJsonStructureFromSlice(mapType)
		if err != nil {
			return "", err
		}
		jsonFormat += structure
	default:
		jsonFormat = fmt.Sprintf("%s", mapType)
	}

	return jsonFormat, nil
}

func GetJsonStructureFromSlice(t reflect.Type) (string, error) {
	var jsonFormat = ""
	var sliceType = t.Elem()

	if sliceType.Kind() == reflect.Struct {
		jsonFormat += "[\n"
		structure, err := GetJsonStructureFromStruct(sliceType)
		if err != nil {
			return "", err
		}
		jsonFormat += structure
		jsonFormat += "\n]"
	} else {
		jsonFormat += fmt.Sprintf("%s", t)
	}

	return jsonFormat, nil
}

func GetJsonStructureFromStruct(t reflect.Type) (string, error) {
	result := fmt.Sprintf("{\n")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Recursively handle nested struct
		if f.Type.Kind() == reflect.Struct {
			if f.Type.Kind() != t.Kind() {
				return "", fmt.Errorf("self referential structs is not allowed")
			}

			subFieldStructure, err := GetJsonStructureFromStruct(f.Type)
			if err != nil {
				return "", err
			}

			result += fmt.Sprintf("    %s: %s,\n", f.Name, subFieldStructure)
		} else {
			result += fmt.Sprintf("    %s: (%s),\n", f.Name, f.Type)
		}
	}

	result += "}"

	return result, nil
}
