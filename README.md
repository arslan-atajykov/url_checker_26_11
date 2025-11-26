# URL Checker Service

Сервис для проверки доступности URL 

## Основные функционалы

### 1. Проверка списка ссылок

`POST /links`

**Request:**

``` json
{
  "links": ["https://google.com", "https://bad.bad"]
}
```

**Response:**

``` json
{
  "links": {
    "https://google.com": "available",
    "https://bad.bad": "not available"
  },
  "links_num": 1
}
```

### 2.  PDF отчёт

`POST /report`

**Request:**

``` json
{
  "links_list": [1, 2]
}
```

**Response:**\
PDF-файл со всеми ссылками и их статусами.

### 3. Хранение задач

Все ссылки сохраняются после завершения программы

------------------------------------------------------------------------

## Запуск

``` bash
go run ./cmd/server
```

--

## Установка зависимостей

Для PDF использовал внешнюю библиотеку:

    github.com/jung-kurt/gofpdf

Установить:

``` bash
go get github.com/jung-kurt/gofpdf
```

------------------------------------------------------------------------

## Примеры запросов через терминал 

### Проверка ссылок 

``` bash
curl -X POST http://localhost:8080/links   -H "Content-Type: application/json"   -d '{"links": ["https://google.com", "https://bad.bad"]}'
```
``` bash
curl -X POST http://localhost:8080/links   -H "Content-Type: application/json"   -d '{"links": ["https://ya.ru", "https://hey.hey"]}'
```
### Генерация PDF

``` bash
curl -X POST http://localhost:8080/report   -H "Content-Type: application/json"   -d '{"links_list":[1,2]}'   --output report.pdf
```
