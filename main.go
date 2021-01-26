package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Upgrade error", 400)
	}
	defer ws.Close()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("receive: %v", msg)
		ws.WriteMessage(msgType, msg)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
