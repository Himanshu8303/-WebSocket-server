package handlers

import (
	"log"
	"net/http"
	"sync"
	"websocket-app/server/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	history []string
	mu      sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		msg := string(message)
		if msg == "history" {
			mu.Lock()
			historyData := utils.JoinStrings(history, "\n")
			mu.Unlock()
			conn.WriteMessage(websocket.TextMessage, []byte(historyData))
		} else {
			reversed := utils.ReverseString(msg)
			conn.WriteMessage(websocket.TextMessage, []byte(reversed))

			mu.Lock()
			history = append(history, msg)
			if len(history) > 5 {
				history = history[len(history)-5:]
			}
			mu.Unlock()
		}
	}
}
