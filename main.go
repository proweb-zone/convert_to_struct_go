package main

import (
	"json_converter/internal/handlers"
)

func main() {

	handlers.CliConverterHandler()

	// jsonString := `{
	//       "name": "Alice",
	//       "details": {
	//           "age": 25,
	//           "job": "Engineer"
	//       }
	//   }`

	// req := map[string]interface{}{}
	// json.Unmarshal([]byte(jsonString), &req)
	// fmt.Println(req)

	// for key, value := range req {
	// 	typeData := reflect.TypeOf(value)
	// 	fmt.Printf("ключ '%s' содержит значение типа %s\n", key, typeData)
	// }

}
