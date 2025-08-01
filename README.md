[![CI](https://github.com/stevenstr/gh-actions-sbs/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/stevenstr/gh-actions-sbs/actions/workflows/ci.yaml)

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

Давай разложим по полочкам основные термины в CI/CD и Docker, чтобы всё стало понятно и логично.

🔄 CI/CD: Непрерывная интеграция и доставка
Термин	Объяснение
CI (Continuous Integration)	Непрерывная интеграция — автоматическая проверка и сборка кода при каждом коммите. Цель — быстро находить ошибки.
CD (Continuous Delivery)	Непрерывная доставка — автоматическая подготовка приложения к выкладке в продакшн. Вы сами решаете, когда выкладывать.
CD (Continuous Deployment)	Непрерывное развертывание — автоматическая выкладка в прод после успешной сборки и тестов. Всё без участия человека.
Pipeline	Последовательность шагов: сборка, тестирование, деплой. Описывается в конфигурационном файле (например, .gitlab-ci.yml).
Stage / Job	Этапы и задачи внутри pipeline. Например: build, test, deploy.
Runner / Agent	Сервер или контейнер, который выполняет pipeline. Например, GitLab Runner.
Artifact	Файл или результат сборки, который передаётся между этапами. Например, бинарник Go.
Trigger	Событие, запускающее pipeline: коммит, merge request, установка тега.

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


# Шаг 5: Обновление тестов До нормального уровня (табличные тесты)
Обновите тесты:

```go
// performRequest создаёт контекст Gin и вызывает handler.
// Возвращает recorder с накопленным ответом.
func performRequest(
	handler gin.HandlerFunc,
	method, path string,
	body io.Reader,
) *httptest.ResponseRecorder {
	// ResponseRecorder будет записывать статус, заголовки и тело
	recorder := httptest.NewRecorder()

	// Получаем пустой Router и Context
	ctx, _ := gin.CreateTestContext(recorder)

	// Формируем HTTP-запрос и помещаем его в ctx
	req := httptest.NewRequest(method, path, body)
	ctx.Request = req

	// Вызываем handler
	handler(ctx)

	return recorder
}
```

Давай разберём эту функцию — performRequest — шаг за шагом, очень просто, как для ребёнка. Представь, что у нас есть крохотный ресторан (наш HTTP-сервер), и мы хотим научиться проверять, что повар (наш хендлер) правильно готовит блюдо (отдаёт правильный ответ). Но чтобы не устраивать полный запуск ресторана, мы делаем «мини-кухню» в тесте. Вот как это выглядит:

### Зачем нужна функция performRequest?
– Мы не хотим каждый раз поднимать настоящий HTTP-сервер, открывать порты и отправлять настоящие запросы.
– Вместо этого мы создаём «двоичный» тест, где загружаем только нужного нам повара (хендлер) и передаём ему искусственный запрос.
– Функция performRequest автоматизирует этот процесс: она готовит среду, даёт хендлеру «запрос» и возвращает нам «ответ», чтобы мы могли его проверить.

###Какие у неё входные и выходные данные?

Вход:
- handler gin.HandlerFunc — это как повар: кусочек кода, который умеет принимать запрос и готовить ответ.
- method и path (строки) — тип запроса (GET, POST и т. д.) и адрес (например, /hello).
- body io.Reader — тело запроса (например, JSON или форма), но в нашем случае мы часто ставим nil, потому что тело нам не нужно.

Выход:
- *httptest.ResponseRecorder — это блокнот и камера, которые записывают всё, что повар отправляет в ответ: какой статус (200, 404…), какие заголовки (Content-Type) и какое тело (JSON-строка).


### Что происходит внутри функции, шаг за шагом: 
#### Шаг 3.1. Создаём «блокнот» для ответа
```go
recorder := httptest.NewRecorder()
```
Представь, что это пустая тетрадка, куда повар будет записывать, что он приготовил: статус, заголовки и тело. 


#### Шаг 3.2. Готовим маленькую «кухню» Gin
```go
ctx, _ := gin.CreateTestContext(recorder)
```

- gin.CreateTestContext даёт нам два животика: пустой маршрутизатор (мы его не используем) и контекст (ctx), в котором хранится и информация о запросе, и «блокнот» для ответа.
- Мы передаём recorder туда, чтобы Gin знал: «Записывай ответ именно в эту тетрадку». 


#### Шаг 3.3. Формируем искусственный HTTP-запрос
```go
req := httptest.NewRequest(method, path, body)
ctx.Request = req
```

- httptest.NewRequest создаёт объект запроса так, будто его прислал клиент (браузер или другая программа).
- method и path определяют, какой запрос — например, GET на /hello.
- body — если нужно, мы можем передать JSON или форму.
- Затем мы говорим нашему контексту ctx, что у него теперь есть запрос: ctx.Request = req. 

#### Шаг 3.4. Вызываем хендлер (повара)
```go
handler(ctx)
```

- Мы передаём контекст с нашим искусственным запросом прямо в функцию-повара.
- Повар прочитает ctx.Request, обработает запрос (прочитает путь, параметры, тело) и напишет результат в ctx.Writer. А ctx.Writer как раз связан с нашим recorder. 

#### Шаг 3.5. Возвращаем записанный ответ
```go
return recorder
```

- В recorder уже записано всё, что хендлер хотел отправить клиенту: например, код 200, заголовок Content-Type: application/json; charset=utf-8 и тело {"text":"Hello, World!"}.
- Мы возвращаем этот recorder из функции, чтобы в тесте проверить:

• Правильный ли статус?

• Правильный ли заголовок Content-Type?

• Правильный ли JSON-ответ?


### Зачем это нужно и где встречается?
- В юнит-тестах HTTP-сервисов на Go с фреймворком Gin (и не только).
- Позволяет очень быстро и локально проверить логику одного конкретного обработчика, без кучи побочных эффектов (нет реальной сети, нет реальных файлов, нет базы данных — только чистый код).
- Если у тебя много мелких хендлеров (регистрация, логин, получение пользователя и т. п.), каждый можно протестировать таким способом.

Итог:
Функция performRequest — это мини-фабрика по созданию искусственного HTTP-запроса и захвату ответа, чтобы мы могли в тестах спокойно и надёжно проверять, как наш хендлер реагирует на разные запросы.


## пример самого теста для хендлера
```go
func TestHelloHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name                string
		method              string
		path                string
		expectedStatusCode  int
		expectedContentType string
		expectedBody        Message
	}{
		{
			name:                "GET hello returns 200 JSON",
			method:              http.MethodGet,
			path:                "/hello",
			expectedStatusCode:  http.StatusOK,
			expectedContentType: "application/json; charset=utf-8",
			expectedBody:        Message{Text: "Hello, World!"},
		},
		// при необходимости можно добавить больше кейсов
	}

	for _, tc := range tests {
		tc := tc // для параллельного запуска
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// выполняем запрос
			recorder := performRequest(HelloHandler, tc.method, tc.path, nil)

			// проверяем статус-код
			require.Equal(t, tc.expectedStatusCode, recorder.Code)

			// проверяем заголовок Content-Type
			require.Equal(t, tc.expectedContentType, recorder.Header().Get("Content-Type"))

			// разбираем JSON-ответ
			var msg Message
			err := json.Unmarshal(recorder.Body.Bytes(), &msg)
			require.NoError(t, err, "response must be valid JSON")

			// проверяем тело ответа
			require.Equal(t, tc.expectedBody, msg)
		})
	}
}

func TestGoodbyeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name                string
		method              string
		path                string
		expectedStatusCode  int
		expectedContentType string
		expectedBody        Message
	}{
		{
			name:                "GET goodbye returns 200 JSON",
			method:              http.MethodGet,
			path:                "/goodbye",
			expectedStatusCode:  http.StatusOK,
			expectedContentType: "application/json; charset=utf-8",
			expectedBody:        Message{Text: "Goodbye, World!"},
		},
		// при необходимости можно добавить больше кейсов
	}

	for _, tc := range tests {
		tc := tc // для параллельного запуска
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// выполняем запрос
			recorder := performRequest(GoodbyeHandler, tc.method, tc.path, nil)

			// проверяем статус-код
			require.Equal(t, tc.expectedStatusCode, recorder.Code)

			// проверяем заголовок Content-Type
			require.Equal(t, tc.expectedContentType, recorder.Header().Get("Content-Type"))

			// разбираем JSON-ответ
			var msg Message
			err := json.Unmarshal(recorder.Body.Bytes(), &msg)
			require.NoError(t, err, "response must be valid JSON")

			// проверяем тело ответа
			require.Equal(t, tc.expectedBody, msg)
		})
	}
}
```

Давай разберём этот тестовый код совершенно по-простому, шаг за шагом, как будто рассказываем новичку или даже ребёнку. Представь, что у нас есть функция (хендлер) HelloHandler, которая на запрос “/hello” отвечает строкой “Hello, World!” в формате JSON. Мы хотим написать автоматический тест, чтобы убедиться, что она всегда так работает.

### Зачем вообще нужны тесты?
- Чтобы без твоего вмешательства проверить, что код делает то, что должно.
- Если что-то случайно поломается, тест сразу об этом сообщит.

## что такое smoke-тест и в чем разница от unit-теста
### 🔥 Smoke-тест (тест "дымовой проверки")
Это базовая проверка, чтобы убедиться, что система вообще запускается и не падает сразу. Название пошло от "проверки на дым": включили устройство — не задымилось, значит, можно тестировать дальше.

- Цель: Быстро удостовериться, что основная функциональность работает.
- Уровень: Обычно проводится после сборки, перед глубоким тестированием.

Примеры:
- Страница загружается и отвечает 200 OK.
- API отдаёт хотя бы базовый ответ.
- Базовые зависимости (БД, Redis, конфиги) доступны.

### 🧪 Unit-тест (модульный тест)
Это подробная проверка маленьких, изолированных частей кода — например, функций, методов или компонентов.

- Цель: Проверить, что конкретный блок логики работает как надо.
- Уровень: Низкий — не зависит от внешней среды.

Примеры:
- Тестирование функции расчёта налога.
- Проверка обработки edge-case данных.
- Убедиться, что метод возвращает нужный результат.


### Какие еще виды тестов существуют?
уществует множество видов тестирования, и каждый из них играет свою роль в обеспечении качества ПО. Вот краткий обзор самых распространённых:

🧪 Unit-тесты
Тестируют отдельные функции или методы.

Изолированы от внешней среды.

Быстрые и точечные.

🔗 Integration-тесты
Проверяют взаимодействие между модулями.

Например, соединение между API и базой данных.

Уловят баги, возникающие при склейке компонентов.

🎭 End-to-End (E2E) тесты
Имитируют реальное поведение пользователя.

Полный путь: от ввода до ответа системы.

Медленные, но эффективные для проверки общего флоу.

💨 Smoke-тесты
Поверхностная проверка, что система вообще работает.

Используются после сборки — чтобы убедиться, что ничего не сломалось.

🧃 Sanity-тесты
Быстрая проверка, что мелкие багфиксы или изменения не нарушили основную функциональность.

Похожи на smoke, но более сфокусированы.

🛠 Regression-тесты
Проверка, что новые изменения не сломали существующий функционал.

Часто автоматизируются и запускаются в CI/CD.

👤 Acceptance-тесты
Оценивают, соответствует ли система требованиям заказчика.

Часто пишутся совместно с бизнесом.

🧼 UI/UX тесты
Проверяют интерфейс и пользовательский опыт.

Может включать визуальную регрессию, удобство взаимодействия и читаемость.

📊 Performance и Load-тесты
Оценивают скорость и устойчивость системы под нагрузкой.

Например, выдерживает ли сайт 1000 запросов в секунду.

#### Сигнатура unit-теста
```go
func TestHelloHandler(t *testing.T) { … }
```
- Любой тест в Go начинается с Test и принимает t *testing.T.
- Инструмент go test найдёт все функции, начинающиеся с Test, и выполнит их.

#### Отключаем лишние логи
```go
gin.SetMode(gin.TestMode)
```
- Gin по умолчанию может много писать в консоль (логи запросов).
- В режиме TestMode он тихонько работает, не мешает выводом.


#### Описываем входные данные и ожидаемые результаты (table-driven тест)
```go
tests := []struct {
  name                string
  method              string
  path                string
  expectedStatusCode  int
  expectedContentType string
  expectedBody        Message
}{
  {
    name:                "GET hello returns 200 JSON",
    method:              http.MethodGet,
    path:                "/hello",
    expectedStatusCode:  http.StatusOK,
    expectedContentType: "application/json; charset=utf-8",
    expectedBody:        Message{Text: "Hello, World!"},
  },
}
```
- tests — это список кейсов (сценариев), которые мы хотим проверить.
- Для каждого кейса указываем:

• name — человекочитаемое название, чтобы понять, что именно проверяется.

• method и path — что мы отправляем хендлеру (GET на /hello).

• expectedStatusCode — какой HTTP-код мы хотим получить (200 OK).

• expectedContentType — какой заголовок “Content-Type” мы хотим увидеть.

• expectedBody — какую структурированную информацию (JSON) мы ждём в теле ответа.


#### Перебираем все кейсы
```go
for _, tc := range tests {
  tc := tc              // фикс для параллели
  t.Run(tc.name, func(t *testing.T) {
    t.Parallel()        // запускаем тесты параллельно
    …
  })
}
```
- for _, tc := range tests значит “для каждого теста из списка”.
- t.Run(tc.name, func(t *testing.T) { … }) создаёт отдельный подпроцесс (subtest) с именем tc.name.
- t.Parallel() позволяет этим подпроцессам выполняться одновременно, ускоряя общий прогон тестов.


#### Внутри каждого подпроцесса делаем три вещи: 
##### 6.1. Выполняем наш хендлер с искусственным запросом
```go
recorder := performRequest(HelloHandler, tc.method, tc.path, nil)   
```   
- performRequest — это вспомогательная функция, которая:

• создаёт фиктивный HTTP-запрос (метод + путь),

• передаёт его в хендлер,

• возвращает recorder — объект, в котором записаны статус, заголовки и тело ответа.

##### 6.2. Проверяем статус-код
```go
require.Equal(t, tc.expectedStatusCode, recorder.Code)      
```

- recorder.Code — это код ответа (например, 200).
- require.Equal (из библиотеки testify) сравнивает два числа и сразу останавливает тест и выводит ошибку, если они не совпадают. 

##### 6.3. Проверяем заголовок Content-Type
```go
require.Equal(t, tc.expectedContentType, recorder.Header().Get("Content-Type")) 
```     
- recorder.Header().Get("Content-Type") читает, какой “Content-Type” хендлер установил.
- Сравниваем с ожидаемым (обычно “application/json; charset=utf-8”). 

##### 6.4. Проверяем тело ответа (JSON)
```go      
  var msg Message      
  err := json.Unmarshal(recorder.Body.Bytes(), &msg)      
  require.NoError(t, err, "response must be valid JSON")      
  require.Equal(t, tc.expectedBody, msg)   
```   
- recorder.Body.Bytes() возвращает «сырые» байты ответа, например {"text":"Hello, World!"}.
- json.Unmarshal пытается распарсить эти байты в переменную msg типа Message (у нас это простая структура с одним полем Text).
- require.NoError проверяет, что распарсилось без ошибок.
- И наконец сравниваем полученную структуру msg с тем, что мы ожидали (tc.expectedBody).


#### Результат
- Если всё совпало (код, заголовок, JSON-тело), тест считается пройденным.
- Если хоть что-то не совпало, require сразу выведет ошибку с подробностями, какой именно чек провалился.

#### Где такое встречается?
- В любом Go-проекте, где есть HTTP-серверы и хочется проверять API-эндоинты без поднятия реального сервера и без реальных сетевых запросов.
- Особенно часто это используется с фреймворком Gin (или похожими), но принцип в целом одинаковый: эмулируем контекст, передаём запрос в хендлер и смотрим, что он “напишет” в ответ.

Таким образом, этот тест — это автоматическая проверка, что HelloHandler всегда отдаёт ровно то, что мы от него ожидаем, в любых условиях.


### Что делать со swagger роутом
Покрывать тестами маршрут Swagger — не обязательно и, как правило, нецелесообразно. Вот почему:

🔍 Назначение маршрута
Этот маршрут служит только для отображения UI документации — он не реализует бизнес-логику и не влияет на поведение API.

Платформа Swagger и её обработчик WrapHandler уже протестированы в рамках библиотеки, которую вы используете.

✅ Когда стоит протестировать
Если вы кастомизировали поведение Swagger UI — например, изменили доступность, добавили авторизацию, внедрили динамическую генерацию документации — тогда тесты могут быть полезны.

В случае CI/CD можно проверить, что маршрут существует и отдаёт 200 OK, чтобы убедиться, что документация доступна на деплое.

🧪 Пример простого smoke-теста
Если всё же хочется покрыть базовой проверкой:

```go
func TestSwaggerRouteAvailable(t *testing.T) {
    gin.SetMode(gin.TestMode)

    router := gin.New()
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    ts := httptest.NewServer(router)
    defer ts.Close()

    resp, err := http.Get(ts.URL + "/swagger/index.html")
    require.NoError(t, err)
    require.Equal(t, http.StatusOK, resp.StatusCode)
}
```

🔍 Что делает этот тест
Проверяет, что маршрут Swagger действительно доступен.

Убеждается, что сервер не отдает ошибку (например, 404 или 500).

Не валидирует содержимое, структуру HTML, версию Swagger — только доступность.

🔥 Почему это smoke-тест
Он поверхностный и быстрый — проверка "жива ли документация".

Не углубляется в бизнес-логику или технические детали.

Идеально подходит для проверки после сборки, перед полноценным запуском тестов.

Если бы мы начали валидировать HTML, параметры Swagger конфигурации, или логику кастомного рендеринга — это был бы уже integration-тест или даже UI-тест.

Но в большинстве случаев это — не приоритет. Лучше сфокусироваться на тестах бизнес-логики API, контроллеров, middleware и прочего функционала.

#### Смок по шагам
😊 Представь, что мы хотим проверить: работает ли страничка Swagger, где показывается документация по нашему API. Мы пишем тест, похожий на маленький автоматический инспектор, который откроет эту страницу и убедится, что она не сломалась.

Вот как это работает по шагам — максимально просто:

🧩 Шаг 1: Подготовим игровую площадку
```go
router := gin.New()
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```
➡️ Мы создаём маршруты, как будто запускаем мини-сервер Gin прямо внутри теста.

🧩 Шаг 2: Включаем мини-сервер
```go
ts := httptest.NewServer(router)
defer ts.Close()
```

➡️ Тут мы запускаем тестовый сервер, он настоящий, но временный и работает только для теста. После теста сам выключится.

🧩 Шаг 3: Отправляем запрос как будто из браузера
```go
resp, err := http.Get(ts.URL + "/swagger/index.html")
```

➡️ Это как открыть в браузере адрес: http://тестовый-сервер/swagger/index.html и посмотреть, что вернётся.

🧩 Шаг 4: Проверяем, не сломалось ли
```go
require.NoError(t, err)
require.Equal(t, http.StatusOK, resp.StatusCode)
```
➡️ Мы говорим:

- 💬 “Ошибок не должно быть”
- 🟢 “Страница должна ответить кодом 200 OK, то есть ‘всё хорошо!’”

🧩 Шаг 5: Включаем "режим тишины"
```go
gin.SetMode(gin.TestMode)
```
➡️ Это как нажать кнопку “Тест” — выключаем лишние сообщения в консоль, чтобы они не мешали тесту.

#### 📦 Итого:
- Это smoke-тест — очень простой, быстрый и важный. Он говорит:
- “Swagger жив? Да! Можно двигаться дальше и тестировать остальную систему.”



## Шаг 6: Обновление GitHub Actions

Чтобы swag работал в GitHub Actions, его нужно установить вручную в рамках workflow. Вот как это делается:
📌 Что происходит:

- go install — скачивает и ставит swag CLI.
- $GITHUB_PATH — добавляет путь к исполняемым файлам Go (~/go/bin) в PATH, чтобы следующая команда могла использовать swag.

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
        go-version: '1.23'

    - name: Install Swagger
      run: |
          go install github.com/swaggo/swag/cmd/swag@v1.8.12
          echo "$(go env GOPATH)/bin" >> $GITHUB_PATH

    - name: Swag version
      run: swag --version

    - name: Generate Swagger docs
      run: swag init

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
```
## Заключение
Теперь у вас есть простой REST API на Go с парой хендлеров, минимальным набором тестов (юнит табличные и смок), интеграцией Swagger для документации и настроенным GitHub Actions для CI/CD. Вы можете расширять этот пример, добавляя больше функциональности, тестов и улучшений по мере необходимости.


# Часть 3
давай теперь добавим docker к этому проекту

🐳 Docker — это как чемодан с заранее собранной средой, в которой твой Go-приложение живёт и работает, независимо от того, где его запустят. Вот простое объяснение:

#### 🧩 Что такое Docker?
Это инструмент, позволяющий упаковать приложение и всё, что ему нужно — зависимости, конфиги, версии — в один контейнер.

Контейнер — как мини-компьютер с твоим приложением внутри. Его можно запускать везде одинаково: на сервере, локально, в облаке.

#### основные термины и понятия Docker
Docker: Контейнеризация приложений
Термин	Объяснение
Docker	Платформа для упаковки приложения и его окружения в контейнер.
Контейнер	Изолированная среда, где работает приложение. Лёгкий, быстрый, переносимый.
Образ (Image)	Шаблон для создания контейнера. Содержит всё необходимое: код, зависимости, конфиги.
Dockerfile	Скрипт, описывающий, как собрать образ.
Volume	Хранилище данных, которое можно подключить к контейнеру.
Network	Виртуальная сеть для связи контейнеров между собой.
Registry	Хранилище образов. Например, Docker Hub или GitLab Container Registry.
Tag	Метка версии образа. Например, myapp:1.0.3.
Compose	Инструмент для запуска нескольких контейнеров вместе (docker-compose.yml).
Multi-stage build	Сборка образа в несколько этапов, чтобы финальный образ был лёгким и безопасным.

#### ⚙️ Зачем он нужен Go-разработчику?
✅ 1. Гарантированное окружение
Go хорошо компилируется в бинарники, но окружение (переменные, версии библиотек, хранилища) всё равно важно.

Docker гарантирует, что у всех участников команды и у прод-серверов будет одинаковая среда.

🚀 2. Простой деплой
С Docker можно задеплоить приложение буквально одной командой: docker run.

Не нужно вручную настраивать сервер, устанавливать Go, пакеты, зависимости.

🔁 3. Лёгкое масштабирование
Каждый Go-сервис можно упаковать в отдельный контейнер.

Управлять контейнерами удобно через Docker Compose, Kubernetes и CI/CD пайплайны.

🧪 4. Автоматизация тестов и сборки
Можно настроить pipeline, который в CI билдит, тестит и выкатывает образ.

Это упрощает поддержку и ускоряет выпуск новых версий.

📦 В чём кайф для Go?
Go — это язык, заточенный под микросервисы и высокую скорость сборки



#### окей, расскажи зачем добавлять docker к микросервису на go
🚢 Отличный вопрос — Docker и микросервисы — как хлеб и масло, особенно в Go-проектах. Объясню, зачем вообще «докеризовать» микросервис:

🧰 1. Изоляция окружения
Docker упаковывает всё: бинарник Go, зависимости, конфиги — в контейнер, работающий одинаково на любой машине.

Забудь про «на моём компе работает, а на сервере нет» — в контейнере будет всегда то же самое окружение.

🚀 2. Лёгкий деплой
Контейнер можно запустить в любой среде: локально, на сервере, в Kubernetes, в облаке.

Никаких ручных установок Go или зависимостей — только docker run.

🔄 3. Масштабируемость
Один микросервис = один контейнер. Нужно больше инстансов? Просто запускаем больше копий.

Это удобно, особенно при использовании оркестраторов: Docker Compose, Kubernetes, Nomad.

🧪 4. Единый pipeline для CI/CD
В CI легко прогнать тесты, собрать образ, опубликовать его, и задеплоить — всё через один Dockerfile.

Работает хорошо с GitHub Actions, GitLab CI и другими инструментами.

🛡 5. Безопасность и стабильность
Контейнер — как песочница. Он изолирован от хост-системы, что снижает риск случайных конфликтов или уязвимостей.

Можно задать точную версию Go, Linux, нужных пакетов.

🔮 Бонус: проще работать в команде
У каждого разработчика одинаковая среда, независимо от ОС.

Даже новичок может просто запустить docker compose up и работать.

Хочешь — могу накидать Dockerfile под типичный микросервис на Gin с Swagger и ENV переменными. Или показать, как собрать multi-stage образ, чтобы он был компактным и безопасным. Какой вариант тебе интереснее? 🛠️

Давайте добавим Docker в наш проект. Мы создадим Dockerfile для контейнеризации нашего приложения и обновим GitHub Actions для сборки и тестирования в Docker контейнере.

## Шаг 1: Создание Dockerfile
Создайте файл Dockerfile в корневой директории проекта и добавьте следующий код:
```sh
# Используем официальный образ Go как базовый образ
FROM golang:1.23-alpine

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


собираем образ:
```sh
docker build -t gh-actions-sbs .
```

запускаем контейнер по имени:
```sh
docker run gh-actions-sbs:latest
```

Чтобы запустить Docker-контейнер на конкретном порте, нужно использовать флаг -p в команде docker run. Он позволяет сопоставить порт хоста с портом внутри контейнера.

🔧 Синтаксис команды
```sh
docker run -p <порт_хоста>:<порт_контейнера> <имя_образа>
```
Например, если твое приложение внутри контейнера слушает порт 8080, а ты хочешь, чтобы оно было доступно на порту 3000 на твоем компьютере:

```sh
docker run -p 3000:5000 gh-actions-sbs:latest
```

3000 — порт на твоем компьютере (хосте)

8080 — порт внутри контейнера, на котором работает приложение

Запуск тестов в контейнере:
```sh
docker run --rm gh-actions-sbs go test -v ./...
```

#### 😶 Если браузер «молчит» и не видит сайт из Docker-контейнера 
— значит, где-то нарушена цепочка: приложение → контейнер → порт → хост → браузер. Давай разберёмся по шагам:

🔍 1. Приложение слушает правильный адрес?
В Go-приложении должно быть:

go
http.ListenAndServe("0.0.0.0:8080", nil)
❗️Если стоит "127.0.0.1" — контейнер не отдаёт наружу, и браузер не сможет достучаться.

📦 2. Порт проброшен при запуске?
Ты должен запускать контейнер с флагом -p, например:

bash
docker run -p 8080:8080 myapp
Первый 8080 — порт на твоём компьютере

Второй 8080 — порт внутри контейнера

🧱 3. Проверка: контейнер работает?
Выполни:

bash
docker ps
Убедись, что контейнер запущен и порт проброшен.

🌐 4. Открытие в браузере
Открой:

http://localhost:8080
Если ты в WSL — попробуй:

http://127.0.0.1:8080
Или IP адрес WSL:

bash
ip addr | grep inet
🧪 5. Быстрая диагностика
В терминале WSL:

bash
curl localhost:8080
Если получаешь ответ — значит, приложение работает, и проблема в доступе из браузера.

🧠 Частые причины:
Приложение слушает 127.0.0.1, а не 0.0.0.0

Порт не проброшен (-p)

Фаервол или антивирус блокирует порт

Приложение упало внутри контейнера — проверь логи:

bash
docker logs <container_id>

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
        go-version: '1.23'

    - name: Install Swag
      run: go install github.com/swaggo/swag/cmd/swag@v1.8.12
          echo "$(go env GOPATH)/bin" >> $GITHUB_PATH

    - name: Swag version
      run: swag --version

    - name: Generate Swagger docs
      run: swag init

    - name: Build Docker image
      run: docker build -t go-rest-api .

    - name: Run tests in Docker container
      run: docker run --rm go-rest-api go test -v ./...

    - name: Run application in Docker container
      run: timeout 15s docker run --rm -p 8080:8080 go-rest-api
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


## graceful shutdown при падении healthcheck.

По-сеньёрному — значит не просто «запустилось», а гарантированно работает, и если нет — джоба падает. Вот несколько подходов, которые используют опытные разработчики Go-проектов в GitHub Actions:

🧰 4. Используй готовый GitHub Action
Например: zethuman/healthcheck

yaml
- name: Container healthcheck
  uses: zethuman/healthcheck@v0.0.1
  with:
    name: my-container
    timeout: 180
    interval: 2
Проверяет статус контейнера и завершает джобу, если он unhealthy

🧼 5. В Go-приложении — отдельный /health эндпоинт
go
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
})
Это стандарт для Kubernetes, Docker и CI/CD — не игнорируй


 Если healthcheck падает — это сигнал, что приложение не в порядке, и его стоит корректно завершить. Вот как можно реализовать graceful shutdown в Gin при сбое healthcheck:

🧠 Общая идея
Периодически проверяем состояние (например, SQL, Redis, ENV).

Если проверка не проходит — инициируем shutdown.

Завершаем сервер через http.Server.Shutdown() с таймаутом.

Пример:
```go
func checkDependencies() bool {
	// проверкИ коннектов к бд редиске и тп

	return true
}
```

#### main c health check и graceful:
```go
func main() {
	// 1. Инициализируем Gin с дефолтными middleware (Logger, Recovery)
	router := gin.Default()

	// 2. Регистрируем любые эндпоинты
	// Healthcheck endpoint
	router.GET("/health", func(c *gin.Context) {
		if !checkDependencies() {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/hello", HelloHandler)
	router.GET("/goodbye", GoodbyeHandler)

	// Роут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 3. Оборачиваем router в http.Server
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	// 4. Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("🚀 Starting server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Канал для внутреннего shutdown при падении healthcheck
	internalShutdown := make(chan struct{})

	// Мониторинг состояния
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if !checkDependencies() {
				log.Println("Healthcheck failed — initiating shutdown")
				internalShutdown <- struct{}{}
				return
			}
		}
	}()

	// 5. Ловим системные сигналы для graceful-shutdown
	// Настраиваем ловлю сигнала прерывания (Ctrl+C / kill)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		log.Println("Получен сигнал завершения")
	case <-internalShutdown:
		log.Println("Healthcheck упал — graceful shutdown")
	}
	log.Println("🔌 Shutdown signal received, exiting...")

	// 6. Останавливаем сервер с таймаутом (пока не обрывать запросы)
	// Даем серверу 5 секунд на «тихую» остановку
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
```


## встраиваем health check в CI
```yaml
    - name: Run App in Docker Container
      run: docker run -d -p 8080:8080 gh-actions-sbs

    - name: Health check (to 20s)
      run: |
        for i in {1..20}; do
          if curl --fail http://localhost:8080/health; then
            echo "health - ok."
            exit 0
          fi
          echo "Loading..."
          sleep 1
        done
        echo "health - bad."
        exit 1
```

🔍 Как это работает
CI запускает приложение в фоне (&)

Ждёт пару секунд (sleep 2)

Делает curl на /health

Если checkDependencies() вернёт false, /health отдаст 503, и curl завершится с ошибкой → CI упадёт

### Добавь статус Ci в ридми

зауди в Actions
выбери любой воркфлоу
создай бейдж
всё
[![CI](https://github.com/stevenstr/gh-actions-sbs/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/stevenstr/gh-actions-sbs/actions/workflows/ci.yaml)

## Переезд на multy-stage building

📦 Что такое multi-stage образ?
Это когда ты собираешь Docker-образ в несколько этапов. Каждый этап — это отдельная стадия, где ты что-то делаешь (сборка, копирование файлов, настройка). И в финальном образе ты берёшь только нужное, остальное — выбрасываешь.

🏋️‍♂️ Зачем это нужно?
🚫 Без multi-stage:
Образ включает всё подряд: компиляторы, лишние пакеты, кеши.

Он тяжёлый и небезопасный — можно случайно оставить утечки или инструменты, которые не должны быть в проде.

✅ С multi-stage:
Образ — минималистичный: только бинарник Go и нужные ENV.

Безопаснее, легче, быстрее грузится и запускается.


💡 Multi-stage образы — мощная техника, но у неё есть свои минусы. Давай разберёмся без фанатизма:

⚠️ 1. Сложнее отладка
Каждый этап сборки — отдельная среда. Не так просто посмотреть, что пошло не так, особенно если ошибка в промежуточной стадии.

Иногда нужно запускать образ с "builder"-стадии вручную, чтобы проверить сборку.

🧠 2. Увеличение сложности Dockerfile
Становится больше инструкций, больше стадий, копирований и переменных.

Для новичков — может выглядеть как "тёмная магия".

🧪 3. Проблемы с кэшированием
Docker кэширует слои, но при multi-stage один лишний COPY или RUN может сбросить кэш для всего этапа.

Это приводит к длинной сборке, особенно при большом проекте.

🧱 4. Размер builder-образа не всегда игнорируется
Если неправильно настроить, финальный образ может случайно включить лишние файлы.

Нужно аккуратно использовать COPY --from=builder, чтобы не затянуть весь /app.

🧩 5. Ограничения в некоторых CI/CD системах
Некоторые старые CI-инструменты не умеют хорошо работать с multi-stage.

Может быть сложно отлаживать образы на нестандартных раннерах или с кастомными pipeline.

🎯 Но если всё грамотно настроить — плюсы перевешивают. Особенно для прод-сборки с компактным и безопасным образом.


🧙‍♂️ Как это выглядит?
Вот «на пальцах» пример:

Dockerfile
# Этап 1: сборка
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o myservice

# Этап 2: финальный образ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myservice .
CMD ["./myservice"]
📉 Итог:
🚀 Быстрее деплой

📦 Меньше размер

🔐 Меньше атак поверхностей


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
