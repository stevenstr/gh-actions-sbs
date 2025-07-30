package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

func TestSwaggerRouteAvailable(t *testing.T) {
	// смок тест
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/swagger/index.html")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHealthRoute(t *testing.T) {
	// смок тест
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/health", HealthCheckHandler)

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
