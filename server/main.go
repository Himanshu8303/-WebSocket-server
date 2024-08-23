package main

import (
	"log"
	"net/http"
	"websocket-app/server/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
