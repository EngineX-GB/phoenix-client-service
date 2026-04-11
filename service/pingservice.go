package service

import (
	"context"
	"fmt"
	"net/http"
	"phoenix-client-service/ws"
	"time"
)

func StartPinger(serviceName string, serverURL string) (stop func()) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("[INFO] Starting Ping Scheduler")
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				ws.Hub.Broadcast <- []byte(fmt.Sprintf("[%s] Pinger stopped", serviceName))
				fmt.Println("[INFO] [Service : " + serviceName + "] Pinger stopped")
				return

			case <-ticker.C:
				resp, err := http.Get(serverURL + "/ping")
				if err != nil {
					ws.Hub.Broadcast <- []byte(fmt.Sprintf("[%s] Ping error: %v", serviceName, err))
					fmt.Println("[ERROR] [Service : "+serviceName+"] Ping error:", err)
					continue
				}
				resp.Body.Close()
				ws.Hub.Broadcast <- []byte(fmt.Sprintf("[%s] Ping OK", serviceName))
				fmt.Println("[INFO] [Service : " + serviceName + "] Ping OK")
			}
		}
	}()

	return cancel
}
