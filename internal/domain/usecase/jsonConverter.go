package usecase

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/proweb-zone/convert_to_struct_go/internal/domain/model"
)

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
	jsonFile := openFile(c.PathIn)
	model.JsonToStructModel(jsonFile)
}

func (c *Converter) XmlConverter() {
	fmt.Println("конвертируем в структуру xml")
	// получаем xml из файла и отправляем его в модель конвертера
	// подготовленную структуру сохраняем в файл
}

func openFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка открытия файла")
	}
	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	return wr.String()
}
