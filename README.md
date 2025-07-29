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

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```


# Шаг 2: Настройка Swagger

Автодок-комментарии (auto-generated documentation comments) — это специальные комментарии в коде, которые расставляются вручную вами, но по заранее определённому шаблону (синтаксису), понятному инструменту генерации документации (в нашем случае — утилите swag).

Как это работает:

Вы пишете над функцией-хендлером (или над пакетом) «мета-комментарии», начинающиеся с // @….
Утилита swag (команда swag init) сканирует ваш исходный код, собирает эти // @…-строки и превращает их в файл docs/swagger.json (или swagger.yaml) в формате OpenAPI/Swagger.
Swagger UI (к которому вы подключаетеся через gin-swagger.WrapHandler) читает этот JSON/YAML и отображает документацию в браузере.
Пример автодок-комментариев в коде (взятый из предыдущего ответа):

```go

// @Summary      Hello World
// @Description  Returns a hello message
// @ID           helloHandler
// @Produce      json
// @Success      200  {object}  Message
// @Router       /hello [get]
func helloHandler(c *gin.Context) { … }
```
Расшифровка основных меток (@ tags):

@Summary — краткое описание метода.
@Description — подробное описание.
@ID — уникальный идентификатор операции (используется внутри Swagger).
@Accept / @Produce — форматы входящих/исходящих данных (json, xml и т. д.).
@Param — описание параметров (query, body, path и пр.).
@Success / @Failure — описание возможных ответов (код, тип, модель данных).
@Router — путь и HTTP-метод в формате <path> [<method>].
Плюс общие метки над пакетом/файлом: @title, @version, @host, @BasePath и пр.
Откуда «высрать» шаблон автодок-комментариев?

Официальная документация swaggo:
https://github.com/swaggo/swag#declarative-comments-format
Примеры в репозитории gin-swagger:
https://github.com/swaggo/gin-swagger#usage
В сорцах проектов, где уже используется swag: посмотрите, как оформлены // @… в популярных Go-репозиториях.
Как начать:
a) Ставите swag в PATH.
b) В корне проекта запускаете swag init — он автоматически создаст /docs.
c) Дописываете // @…-комментарии к вашим хендлерам (и к корню пакета, чтобы задать общие параметры API).
d) Перезапускаете swag init, обновляете /docs/swagger.json.
e) Открываете UI по /swagger/index.html.

Таким образом «автодок-комментарии» — это просто ваша разметка исходников под генератор Swagger-документации.


Опишем апи в main.go
```go
// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/stevenstr/gh-actions-sbs/docs" // docs генерится swag-ом
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Message — структура ответа
type Message struct {
	Text string `json:"text"`
}

// @title           Simple API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @Summary      Hello World
// @Description  Returns a hello message
// @ID           helloHandler
// @Produce      json
// @Success      200  {object}  Message
// @Router       /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Text: "Hello, World!"})
}

// @Summary      Goodbye World
// @Description  Returns a goodbye message
// @ID           goodbyeHandler
// @Produce      json
// @Success      200  {object}  Message
// @Router       /goodbye [get]
func goodbyeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Text: "Goodbye, World!"})
}

func main() {
	r := gin.Default()

	r.GET("/hello", helloHandler)
	r.GET("/goodbye", goodbyeHandler)

	// Роут для Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
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



# Шаг 5: Запуск и проверка
Ниже разберём пример «старта» HTTP-сервера с возможностью плавного (graceful) завершения по шагам.

```go
import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/gin-gonic/gin"
)
```

– Пакет net/http и github.com/gin-gonic/gin для запуска HTTP-сервера.
– context, os/signal, syscall, time – для организации graceful-shutdown.
– log – для логирования.

```go
func main() {
  // 1. Инициализируем Gin с дефолтными middleware (Logger, Recovery)
	router := gin.Default()

   // 2. Регистрируем любые эндпоинты
	router.GET("/hello", helloHandler)
	router.GET("/goodbye", goodbyeHandler)

	// Роут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  // 3. Оборачиваем router в http.Server
//   – Оборачиваем gin.Engine (реализующий http.Handler) в стандартный http.Server.
// – Благодаря этому можем управлять запуском и завершением более тонко, чем вызывая просто r.Run().
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
 

	// 4. Запускаем сервер в отдельной горутине
//   – Запускаем ListenAndServe() в новой горутине, чтобы основная программа не блокировалась.
// – Если ошибка не равна http.ErrServerClosed (это «нормальный» код закрытия), логируем и завершаем приложение через log.Fatalf.
	go func() {
		log.Printf("🚀 Starting server on %s", srv.Addr)
		if err :=  srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

// 5. Ловим системные сигналы для graceful-shutdown
	// Настраиваем ловлю сигнала прерывания (Ctrl+C / kill)
//   – Создаём канал quit для системных сигналов.
// – signal.Notify подписывается на SIGINT (Ctrl+C) и SIGTERM (kill).
// – <-quit блокируется до получения одного из этих сигналов.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🔌 Shutdown signal received, exiting...")

// 6. Останавливаем сервер с таймаутом (пока не обрывать запросы)
	// Даем серверу 5 секунд на «тихую» остановку
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
//   – Создаём контекст с таймаутом (5 секунд), чтобы не ждать бесконечно.
// – srv.Shutdown(ctx) сообщит серверу:
// • перестать принимать новые соединения;
// • дать «живущим» запросам завершиться в течение времени таймаута;
// • после чего принудительно закрыть остатки.
// – Если в пределах 5 секунд все запросы не завершились — Shutdown вернёт ошибку, и мы логируем фатал.
// – Если всё прошло успешно — пишем в лог об удачном завершении.

	log.Println("🛑 Server stopped gracefully")
}
```

Пояснения шагов:
- gin.Default() — заводит стандартный HTTP-сервер с логированием и recover-middleware.
- Роуты регистрируются привычным router.GET/....
- Создаётся стандартный http.Server, в поле Handler передаётся наш Gin-router.
srv.ListenAndServe() запускается в отдельной горутине, чтобы основной поток мог «сидеть» и ждать сигнала остановки.
- Через os.Signal и signal.Notify ловим Ctrl+C (SIGINT) или kill (SIGTERM).
- srv.Shutdown(ctx) — это встроенный метод Go-сервера, который: • перестаёт принимать новые подключения;
• даёт текущим обработчикам (Gin-хендлерам) до 5 секунд на завершение;
• затем принудительно закрывает оставшиеся.
Так мы получаем «мягкую» остановку сервера, не «режем» на лету открытые HTTP-соединения.


Зачем так делают?
• Грейсфул-шадоун (graceful shutdown) нужен для того, чтобы при рестарте/обновлении/остановке сервера не обрывать «на лету» активные HTTP-сессии или транзакции.
• Это критично в боевых системах, чтобы клиенты получали ответы, а не «обрыв соединения».
• Использование http.Server + контекст + сигналов OS — стандартный паттерн в Go для управления жизненным циклом HTTP-сервисов.

Что здесь улучшено:
- Убираем избыточный os.Exit(1) после log.Fatal (он и так вызывает os.Exit(1)).
- Пишем понятный log.Fatalf("…: %v", err) вместо errors.Error(err).
- Используем стандартный http.Server для возможности graceful-shutdown.
- Ловим сигналы SIGINT/SIGTERM и даём серверу 5 секунд на корректное завершение активных соединений.


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
