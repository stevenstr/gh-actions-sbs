# gh-actions-sbs

# Go REST API Project

## Описание проекта

Проект представляет собой простой REST API, написанный на языке программирования Go. Основная цель проекта — демонстрация базовых принципов создания REST API, интеграции документации с помощью Swagger, контейнеризации с использованием Docker и автоматизации процессов сборки и тестирования с помощью GitHub Actions.

## Функции проекта

1. **REST API**:
   - Два хендлера: `/hello` и `/goodbye`, которые возвращают JSON-ответы с текстом "Hello, World!" и "Goodbye, World!" соответственно.
   - Использование фреймворка Gin для создания маршрутов и обработки запросов.

2. **Документация с помощью Swagger**:
   - Автоматическая генерация документации API в формате OpenAPI.
   - Доступ к документации через маршрут `/swagger/index.html`.

3. **Контейнеризация с Docker**:
   - Создание Dockerfile для сборки и запуска приложения в контейнере.
   - Упрощение развертывания и обеспечение консистентности окружения.

4. **Автоматизация с GitHub Actions**:
   - Настройка CI/CD pipeline для автоматической сборки, тестирования и запуска приложения при каждом пуше в ветку `main` или при создании pull request.
   - Интеграция с Docker для выполнения тестов и запуска приложения в контейнере.

## Технологии

- **Go**: Язык программирования для создания REST API.
- **Gin**: Веб-фреймворк для Go, используемый для создания маршрутов и обработки запросов.
- **Swagger**: Инструмент для генерации документации API в формате OpenAPI.
- **Docker**: Платформа для контейнеризации приложений.
- **GitHub Actions**: Инструмент для автоматизации процессов CI/CD.

## Механизмы

1. **Создание REST API**:
   - Определение маршрутов и хендлеров с использованием Gin.
   - Обработка HTTP-запросов и возвращение JSON-ответов.

2. **Генерация документации**:
   - Использование аннотаций для описания маршрутов и параметров.
   - Генерация файла `swagger.json` с помощью команды `swag init`.

3. **Контейнеризация**:
   - Создание Dockerfile для сборки приложения.
   - Установка зависимостей и сборка приложения в контейнере.
   - Запуск контейнера с приложением.

4. **Автоматизация CI/CD**:
   - Настройка GitHub Actions для автоматической сборки и тестирования.
   - Использование Docker для выполнения тестов и запуска приложения в контейнере.
   - Автоматическое выполнение workflow при пуше в ветку `main` или при создании pull request.

## Базовая структура после выполгнения частей 1-3
go-rest-api/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── docs/
│   ├── docs.go
│   └── swagger.json
├── tests/
│   └── main_test.go
└── .github/
└── workflows/
└── ci.yml

## Базовая структура после выполнения части 4
go-rest-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   │   └── handler.go
│   ├── config/
│   │   └── config.go
│   └── swagger/
│       └── docs.go
├── pkg/
│   └── utils/
│       └── utils.go
├── docs/
│   └── swagger.json
├── tests/
│   └── main_test.go
├── Dockerfile
├── go.mod
├── go.sum
└── .github/
    └── workflows/
        └── ci.yml


## Установка и запуск

1. **Клонируйте репозиторий**:
```sh
git clone https://github.com/yourusername/go-rest-api.git
cd go-rest-api
```

2. **Установите зависимости:
```sh

go mod tidy
```

3. **Соберите и запустите приложение:
```sh

go run main.go
```

4. **Соберите Docker образ и запустите контейнер:
```sh

docker build -t go-rest-api .
docker run --rm -p 8080:8080 go-rest-api
```

5. **Откройте браузер и перейдите по адресам:
http://localhost:8080/hello
http://localhost:8080/goodbye
http://localhost:8080/swagger/index.html (для документации Swagger)

# Лицензия
Этот проект лицензирован под лицензией MIT.



# Часть 01
Давайте создадим простой REST API на Go с парой хендлеров и без базы данных. Затем мы добавим минимальный набор тестов и настроим GitHub Actions для CI/CD без пуша в Docker Hub.

## Шаг 1: Создание проекта
Создайте новую директорию для вашего проекта и перейдите в нее:
```sh
mkdir go-rest-api
cd go-rest-api
```

Инициализируйте новый модуль Go:
```sh
go mod init go-rest-api
```

## Шаг 2: Создание REST API

Создайте файл main.go и добавьте следующий код:
```go
package main

import (
    "encoding/json"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Hello, World!"}
    json.NewEncoder(w).Encode(message)
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Goodbye, World!"}
    json.NewEncoder(w).Encode(message)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/goodbye", goodbyeHandler)

    http.ListenAndServe(":8080", nil)
}
```


## Шаг 3: Написание тестов
Создайте директорию файл main_test.go:
```sh

 main_test.go
```

Установи пакет:
```sh
go get github.com/stretchr/testify/assert
```

Добавьте следующий код в tests/main_test.go:
```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"text":"Hello GH Actions World!"}`, rr.Body.String())
}

func TestGoodbyeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/goodbye", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GoodbyeHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"text":"Goodbye GH Actions World!"}`, rr.Body.String())
}

```


## Шаг 4: Настройка GitHub Actions
Создайте директорию .github/workflows и файл ci.yml внутри нее:
```sh

mkdir -p .github/workflows
touch .github/workflows/ci.yml
```

### Общая структура
Файл .github/workflows/ci.yml является конфигурационным файлом для GitHub Actions. Он определяет, как и когда должны выполняться автоматические задачи (workflows) в вашем репозитории. В данном случае, это CI (Continuous Integration) pipeline, который будет автоматически запускаться при определенных событиях.

Этот файл автоматизирует процесс сборки и тестирования вашего проекта. Когда вы делаете пуш в ветку main или создаете pull request в эту ветку, GitHub Actions автоматически:

Клонирует ваш репозиторий.
Устанавливает Go.
Компилирует ваш проект.
Запускает все тесты.
Это помогает убедиться, что ваш код работает корректно и не содержит ошибок перед тем, как он будет объединен в основную ветку.

Добавьте следующий код в ci.yml:
```yaml
name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.18'

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
```


### Разбор файла
```yaml
name: CI
```
name: Это имя вашего workflow. В данном случае, оно называется "CI", что означает Continuous Integration.

```yaml
on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
```
on: Определяет события, при которых будет запускаться этот workflow.

push: Workflow будет запускаться при каждом пуше в ветку main.

pull_request: Workflow будет запускаться при каждом создании pull request в ветку main.


```yaml
jobs:
  build:
    runs-on: ubuntu-latest
```
jobs: Определяет набор задач (jobs), которые будут выполняться в этом workflow.

build: Имя задачи. В данном случае, задача называется "build".

runs-on: Определяет операционную систему, на которой будет выполняться задача. В данном случае, это ubuntu-latest, что означает последнюю версию Ubuntu.


```yaml
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
```
steps: Определяет последовательность шагов, которые будут выполняться в этой задаче.

name: Описание шага.

uses: Указывает, какой action использовать. В данном случае, это actions/checkout@v2, который клонирует репозиторий в рабочую директорию.


```yaml
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.18'
```
name: Описание шага.

uses: Указывает, какой action использовать. В данном случае, это actions/setup-go@v2, который устанавливает Go на рабочую машину.

with: Передает параметры в action. В данном случае, устанавливается версия Go 1.18.

```yaml
    - name: Build
      run: go build -v ./...
```
name: Описание шага.

run: Команда, которая будет выполнена на этом шаге. В данном случае, это go build -v ./..., которая компилирует все пакеты в проекте.


```yaml
    - name: Test
      run: go test -v ./...
```
name: Описание шага.

run: Команда, которая будет выполнена на этом шаге. В данном случае, это go test -v ./..., которая запускает все тесты в проекте.




## Шаг 5: Запуск и тестирование
Запустите ваше приложение локально:
```sh
go run main.go
```

Откройте браузер и перейдите по адресам http://localhost:8080/hello и http://localhost:8080/goodbye, чтобы убедиться, что API работает.


Запустите тесты локально:
```sh

go test -v ./...
```

## Шаг 6: Запуск GitHub Actions
Создайте репозиторий на GitHub и загрузите туда ваш проект.
GitHub Actions автоматически запустит CI pipeline при каждом пуше в ветку main или при создании  pull request в эту ветку. Вы можете следить за прогрессом в разделе "Actions" вашего репозитория на GitHub.

## Шаг 7: Дополнительные улучшения (опционально)
Если вы хотите добавить дополнительные улучшения, такие как логгирование, middleware или более сложные тесты, вы можете сделать это следующим образом:

### Логгирование:
Добавьте пакет log/slog для логгирования запросов и ответов.
```go
import (
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request at /hello")
    message := Message{Text: "Hello, World!"}
    json.NewEncoder(w).Encode(message)
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request at /goodbye")
    message := Message{Text: "Goodbye, World!"}
    json.NewEncoder(w).Encode(message)
}
```

### Middleware:
Создайте middleware для логгирования времени выполнения запросов.
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("Method: %s, Path: %s, Duration: %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func main() {
    http.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))
    http.Handle("/goodbye", loggingMiddleware(http.HandlerFunc(goodbyeHandler)))

    http.ListenAndServe(":8080", nil)
}
```


## Более сложные тесты:
Добавьте тесты для проверки middleware и других аспектов вашего приложения.
```go
func TestLoggingMiddleware(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := loggingMiddleware(http.HandlerFunc(helloHandler))
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, `{"text":"Hello, World!"}`, rr.Body.String())
}
```


## Заключение
Теперь у вас есть простой REST API на Go с парой хендлеров, минимальным набором тестов и настроенным GitHub Actions для CI/CD. Вы можете расширять этот пример, добавляя больше функциональности, тестов и улучшений по мере необходимости.



# Часть 2
давай добавим к этому проекту swagger

Swagger (OpenAPI) — это мощный инструмент для документации API. Давайте добавим Swagger в наш проект на Go. Мы будем использовать библиотеку swaggo и swaggo/swag для генерации документации Swagger.

## Шаг 1: Установка зависимостей
Установите необходимые пакеты:
```sh
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```


# Шаг 2: Настройка Swagger
Создайте файл docs/docs.go для настройки Swagger:
```go
// docs/docs.go
package docs

import (
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url http://swagger.io

func SwaggerInit(router *gin.Engine) {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

Обновите файл main.go для интеграции Swagger:
```go

package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "go-rest-api/docs"
)

type Message struct {
    Text string `json:"text"`
}

// @Summary Hello World
// @Description do ping
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} Message
// @Router /hello [get]
func helloHandler(c *gin.Context) {
    message := Message{Text: "Hello, World!"}
    c.JSON(http.StatusOK, message)
}

// @Summary Goodbye World
// @Description do ping
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} Message
// @Router /goodbye [get]
func goodbyeHandler(c *gin.Context) {
    message := Message{Text: "Goodbye, World!"}
    c.JSON(http.StatusOK, message)
}

func main() {
    router := gin.Default()

    router.GET("/hello", helloHandler)
    router.GET("/goodbye", goodbyeHandler)

    docs.SwaggerInit(router)

    router.Run(":8080")
}
```

## Шаг 3: Генерация документации
Сгенерируйте документацию Swagger:
```sh

swag init
```
Эта команда создаст файл docs/swagger.json, который содержит документацию API в формате OpenAPI.

# Шаг 4: Запуск и проверка
Запустите ваше приложение:
```sh

go run main.go
```

Откройте браузер и перейдите по адресу http://localhost:8080/swagger/index.html, чтобы увидеть документацию Swagger.


# Шаг 5: Обновление тестов
Обновите тесты для проверки маршрутов Swagger:
```go

package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
    "go-rest-api/docs"
    "go-rest-api/main"
)

func TestHelloHandler(t *testing.T) {
    router := gin.Default()
    router.GET("/hello", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, `{"text":"Hello, World!"}`, rr.Body.String())
}

func TestGoodbyeHandler(t *testing.T) {
    router := gin.Default()
    router.GET("/goodbye", main.GoodbyeHandler)
    docs.SwaggerInit(router)

    req, err := http.NewRequest("GET", "/goodbye", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, `{"text":"Goodbye, World!"}`, rr.Body.String())
}

func TestSwaggerHandler(t *testing.T) {
    router := gin.Default()
    docs.SwaggerInit(router)

    req, err := http.NewRequest("GET", "/swagger/index.html", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
}
```

## Шаг 6: Обновление GitHub Actions
Обновите файл .github/workflows/ci.yml, чтобы включить генерацию документации Swagger:
```yaml

name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.18'

    - name: Install Swag
      run: go get -u github.com/swaggo/swag/cmd/swag

    - name: Generate Swagger docs
      run: swag init

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
```
## Заключение
Теперь у вас есть простой REST API на Go с парой хендлеров, минимальным набором тестов, интеграцией Swagger для документации и настроенным GitHub Actions для CI/CD. Вы можете расширять этот пример, добавляя больше функциональности, тестов и улучшений по мере необходимости.


# Часть 3
давай теперь добавим docker к этому проекту

Давайте добавим Docker в наш проект. Мы создадим Dockerfile для контейнеризации нашего приложения и обновим GitHub Actions для сборки и тестирования в Docker контейнере.

## Шаг 1: Создание Dockerfile
Создайте файл Dockerfile в корневой директории проекта и добавьте следующий код:
```sh
# Используем официальный образ Go как базовый образ
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN go build -o main .

# Указываем, что контейнер будет запускать собранное приложение
CMD ["./main"]
```


## Шаг 2: Обновление GitHub Actions
Обновите файл .github/workflows/ci.yml, чтобы включить сборку и тестирование в Docker контейнере:
```yaml

name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.18'

    - name: Install Swag
      run: go get -u github.com/swaggo/swag/cmd/swag

    - name: Generate Swagger docs
      run: swag init

    - name: Build Docker image
      run: docker build -t go-rest-api .

    - name: Run tests in Docker container
      run: docker run --rm go-rest-api go test -v ./...

    - name: Run application in Docker container
      run: docker run --rm -p 8080:8080 go-rest-api
```

Шаг 3: Запуск и проверка локально
Соберите Docker образ локально:
```sh

docker build -t go-rest-api .
```

Запустите контейнер:
```sh
docker run --rm -p 8080:8080 go-rest-api
```

Откройте браузер и перейдите по адресам http://localhost:8080/hello и http://localhost:8080/goodbye, чтобы убедиться, что API работает.

Откройте браузер и перейдите по адресу http://localhost:8080/swagger/index.html, чтобы увидеть документацию Swagger.

## Заключение
Теперь у вас есть простой REST API на Go с парой хендлеров, минимальным набором тестов, интеграцией Swagger для документации, Docker контейнеризацией и настроенным GitHub Actions для CI/CD. Вы можете расширять этот пример, добавляя больше функциональности, тестов и улучшений по мере необходимости.


# Часть 4
оформи теперь этот проект в классическом гошном стиле с папками cmd, internal и прочее

Давайте реорганизуем проект в классическом Go стиле, используя папки cmd, internal, pkg и docs. Это поможет лучше структурировать код и сделать его более модульным и поддерживаемым.

go-rest-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   │   └── handler.go
│   ├── config/
│   │   └── config.go
│   └── swagger/
│       └── docs.go
├── pkg/
│   └── utils/
│       └── utils.go
├── docs/
│   └── swagger.json
├── tests/
│   └── main_test.go
├── Dockerfile
├── go.mod
├── go.sum
└── .github/
    └── workflows/
        └── ci.yml


Содержимое файлов
cmd/main.go

package main

import (
    "github.com/gin-gonic/gin"
    "go-rest-api/internal/api"
    "go-rest-api/internal/swagger"
)

func main() {
    router := gin.Default()

    api.RegisterHandlers(router)
    swagger.SwaggerInit(router)

    router.Run(":8080")
}


Содержимое файлов
internal/api/handler.go

package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Message struct {
    Text string `json:"text"`
}

func HelloHandler(c *gin.Context) {
    message := Message{Text: "Hello, World!"}
    c.JSON(http.StatusOK, message)
}

func GoodbyeHandler(c *gin.Context) {
    message := Message{Text: "Goodbye, World!"}
    c.JSON(http.StatusOK, message)
}

func RegisterHandlers(router *gin.Engine) {
    router.GET("/hello", HelloHandler)
    router.GET("/goodbye", GoodbyeHandler)
}

Содержимое файлов
internal/swagger/docs.go
package swagger

import (
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url http://swagger.io

func SwaggerInit(router *gin.Engine) {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}



Содержимое файлов
tests/main_test.go


Содержимое файлов
Содержимое файлов
Содержимое файлов
