package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved request at /hello")
	message := Message{"Hello GH Actions World!"}
	json.NewEncoder(w).Encode(message)
}
func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved request at /goodbye")
	message := Message{"Goodbye GH Actions World!"}
	json.NewEncoder(w).Encode(message)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Method: %s, Path: %s, Duration: %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(HelloHandler)))

	mux.Handle("/goodbye", loggingMiddleware(http.HandlerFunc(GoodbyeHandler)))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}
