package usecase

import "fmt"

type Converter struct {
	PathIn  string
	PathOut string
}

type IConverter interface {
	JsonConverter()
	XmlConverter()
}

func InitJSonConverter(pathIn string, pathOut string) *Converter {
	return &Converter{PathIn: pathIn, PathOut: pathOut}
}

func (c *Converter) JsonConverter() {
	fmt.Println("конвертируем в структуру json")
	// получаем json из файла и отправляем его в модель конвертера
	// подготовленную структуру сохраняем в файл
}

func (c *Converter) XmlConverter() {
	fmt.Println("конвертируем в структуру xml")
	// получаем xml из файла и отправляем его в модель конвертера
	// подготовленную структуру сохраняем в файл
}
