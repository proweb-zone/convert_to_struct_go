package usecase

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/proweb-zone/convert_to_struct_go/internal/domain/model"
)

type Converter struct {
	PathIn  string
	PathOut string
}

type IConverter interface {
	JsonConverter()
}

func InitJSonConverter(pathIn string, pathOut string) *Converter {
	return &Converter{PathIn: pathIn, PathOut: pathOut}
}

func (c *Converter) JsonConverter() {
	fmt.Println("конвертируем в структуру json")
	jsonFile, err := openFile(c.PathIn)

	if err != nil {
		fmt.Errorf("Error open file")
	}

	data, err := model.JsonToStructModel(jsonFile)

	if err != nil {
		fmt.Errorf("Error: Convert json to struct " + err.Error())
		return
	}

	writeStructToFile(data, c.PathOut)

}

func openFile(path string) (output string, err error) {
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

	return wr.String(), nil
}

func writeStructToFile(data string, path string) error {

	// Создаем все несуществующие дирректории на пути
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Errorf("Error create directory" + err.Error())
		return err
	}

	// Создаем файл
	file, err := os.Create(path + "/" + "testEntity.go")
	if err != nil {
		fmt.Errorf("Error create file " + err.Error())
		return err
	}
	defer file.Close()

	packageName := packageName(path)

	// Записать текст в файл
	_, err = file.WriteString(packageName)
	if err != nil {
		fmt.Errorf("Error write structure to file " + err.Error())
		return err
	}

	// Сохранить изменения в файле
	err = file.Sync()
	if err != nil {
		fmt.Errorf("Error save file" + err.Error())
		return err
	}

	return nil
}

func packageName(path string) string {
	pathList := strings.Split(path, "/")

	// Создать новый массив без пустых элементов
	var newPathList []string
	for _, value := range pathList {
		if value != "" {
			newPathList = append(newPathList, value)
		}
	}

	lenElemPath := len(newPathList)

	// создаем буфер для хранения package name
	bufPackageName := bytes.NewBufferString("")
	bufPackageName.WriteString("package ")

	// если корневая дирректория проекта
	if newPathList[0] == "." && lenElemPath == 1 {
		bufPackageName.WriteString("main")
	} else {
		bufPackageName.WriteString(newPathList[lenElemPath-1])
	}

	packageName := bufPackageName.String()
	return packageName
}
