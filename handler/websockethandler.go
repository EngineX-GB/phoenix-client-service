package handler

import (
	"fmt"
	"net/http"
	"phoenix-client-service/ws"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(res http.ResponseWriter, req *http.Request) {
	conn, err := Upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("Error Upgrading:", err)
		return
	}

	// Register client
	ws.Hub.Mu.Lock()
	ws.Hub.Clients[conn] = true
	ws.Hub.Mu.Unlock()

	fmt.Println("[DEBUG] Client connected")

	// Listen for incoming messages (optional)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("[DEBUG] Client disconnected:", err)

			ws.Hub.Mu.Lock()
			delete(ws.Hub.Clients, conn)
			ws.Hub.Mu.Unlock()

			conn.Close()
			break
		}

		fmt.Println("[DEBUG] Received:", string(msg))
	}

}
