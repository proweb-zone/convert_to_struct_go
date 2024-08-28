package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonToStructModel(data string) (buildStruct string, err error) {
	prepairData(data)
	return "", nil
}

func prepairData(data string) {
	mapData := map[string]interface{}{}

	json.Unmarshal([]byte(data), &mapData)

	for key, value := range mapData {
		typeData := reflect.TypeOf(value)
		if typeData == reflect.TypeOf(map[string]interface{}(nil)) {
			json, _ := json.Marshal(value)
			prepairData(string(json))
		} else {
			fmt.Printf("ключ '%s' содержит значение типа %s\n", key, typeData)
		}
	}
}
