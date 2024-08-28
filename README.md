# Библиотека для преобразования  JSON в структуру на Go

### Утановка

```golang
git clone https://github.com/proweb-zone/convert_to_struct_go
```
### Сборка

```golang
go build .
```

### Example

```golang
go run . -json ./mockData/user.json ./resultData
```

Аргумент 1 - передается тип данных '-json'

Аргумент 2 - json файл для конвертации

Аргумент 3 - путь для сохрания результата
