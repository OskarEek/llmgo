package jsonhelper

import (
	"fmt"
	"reflect"
)

func GetJsonStructureFromType(mapType reflect.Type) (string, error) {
	mapType = TransformFromPointer(mapType)
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

func TransformFromPointer(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func GetJsonStructureFromSlice(t reflect.Type) (string, error) {
	t = TransformFromPointer(t)
	var jsonFormat = ""
	var sliceType = TransformFromPointer(t.Elem())

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
	t = TransformFromPointer(t)
	result := fmt.Sprintf("{\n")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldType := TransformFromPointer(f.Type)

		// Recursively handle nested struct
		if fieldType.Kind() == reflect.Struct {

			//TODO: figure out how to solve 2 structs referencing eachother since this will only work if a struct references itself
			//Prevent self referential structs
			if fieldType.Name() == t.Name() {
				return "", fmt.Errorf("self referential structs is not allowed")
			}

			subFieldStructure, err := GetJsonStructureFromStruct(fieldType)
			if err != nil {
				return "", err
			}

			result += fmt.Sprintf("    %s: %s,\n", f.Name, subFieldStructure)
		} else {
			result += fmt.Sprintf("    %s: (%s),\n", f.Name, fieldType)
		}
	}

	result += "}"

	return result, nil
}
