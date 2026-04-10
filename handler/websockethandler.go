package handler

import (
	"fmt"
	"net/http"

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
	defer conn.Close()

	// listen to incoming messages
	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in reading message", err)
			break
		}
		fmt.Println("[INFO] Received message : %s\\n", string(message))

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message", err)
			break
		}
	}

}
