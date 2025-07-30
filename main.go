// main.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Text: "Hello, World!"})
}

// @Summary      Goodbye World
// @Description  Returns a goodbye message
// @ID           goodbyeHandler
// @Produce      json
// @Success      200  {object}  Message
// @Router       /goodbye [get]
func GoodbyeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Text: "Goodbye, World!"})
}

func main() {
	// 1. Инициализируем Gin с дефолтными middleware (Logger, Recovery)
	router := gin.Default()

	// 2. Регистрируем любые эндпоинты
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

	// 5. Ловим системные сигналы для graceful-shutdown
	// Настраиваем ловлю сигнала прерывания (Ctrl+C / kill)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🔌 Shutdown signal received, exiting...")

	// 6. Останавливаем сервер с таймаутом (пока не обрывать запросы)
	// Даем серверу 5 секунд на «тихую» остановку
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("🛑 Server stopped gracefully")
}
