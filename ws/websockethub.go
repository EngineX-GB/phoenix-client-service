package ws

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan []byte
	Mu        sync.Mutex
}

var Hub = WebSocketHub{
	Clients:   make(map[*websocket.Conn]bool),
	Broadcast: make(chan []byte),
}

func (h *WebSocketHub) Run() {
	for {
		msg := <-h.Broadcast
		fmt.Printf("Broadcast message : %s\n", msg)

		h.Mu.Lock()
		for client := range h.Clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error writing to client:", err)
				client.Close()
				delete(h.Clients, client)
			}
		}
		h.Mu.Unlock()
	}
}
