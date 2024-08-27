package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonToStructModel(jsonData string) {

	req := map[string]interface{}{}
	json.Unmarshal([]byte(jsonData), &req)

	for key, value := range req {
		typeData := reflect.TypeOf(value)
		if typeData == reflect.TypeOf(map[string]interface{}(nil)) {
			prepairData(value)
		} else {
			fmt.Printf("ключ '%s' содержит значение типа %s\n", key, typeData)
		}
	}

}

func prepairData(data interface{}) {

	req := map[string]interface{}{}

	jsonBytes, _ := json.Marshal(data)
	json.Unmarshal([]byte(jsonBytes), &req)

	for key, value := range req {
		typeData := reflect.TypeOf(value)
		if typeData == reflect.TypeOf(map[string]interface{}(nil)) {
			prepairData(value)
		} else {
			fmt.Printf("ключ '%s' содержит значение типа %s\n", key, typeData)
		}
	}
}
