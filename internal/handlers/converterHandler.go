package handlers

import (
	"fmt"
	"json_converter/internal/domain/usecase"
	"os"
)

func CliConverterHandler() {

	if len(os.Args) < 3 {
		fmt.Println("Error: There are not enough arguments")
		return
	}

	cliArguments := os.Args

	converter := usecase.InitJSonConverter(cliArguments[2], cliArguments[3])

	switch cliArguments[1] {
	case "-json":
		converter.JsonConverter()
	case "-xml":
		converter.XmlConverter()
	}
}
