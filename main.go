package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/goodbye", GoodbyeHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}
