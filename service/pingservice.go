package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"phoenix-client-service/model"
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
				msg := model.Notification{
					NotificationType: "ping",
					Content: model.Content{
						Operation:   "ping",
						ServiceName: serviceName,
						Status:      "STOPPED",
						Detail:      "STOPPED",
					},
				}
				jsonBytes, _ := json.Marshal(msg)
				ws.Hub.Broadcast <- jsonBytes

				fmt.Println("[INFO] [Service : " + serviceName + "] Pinger stopped")
				return

			case <-ticker.C:
				resp, err := http.Get(serverURL + "/ping")
				if err != nil {

					msg := model.Notification{
						NotificationType: "ping",
						Content: model.Content{
							Operation:   "ping",
							ServiceName: serviceName,
							Status:      "ERROR",
							Detail:      err.Error(),
						},
					}
					jsonBytes, _ := json.Marshal(msg)
					ws.Hub.Broadcast <- jsonBytes

					fmt.Println("[ERROR] [Service : "+serviceName+"] Ping error:", err)
					continue
				}
				resp.Body.Close()

				msg := model.Notification{
					NotificationType: "ping",
					Content: model.Content{
						Operation:   "ping",
						ServiceName: serviceName,
						Status:      "OK",
						Detail:      "OK",
					},
				}
				jsonBytes, _ := json.Marshal(msg)
				ws.Hub.Broadcast <- jsonBytes

				fmt.Println("[INFO] [Service : " + serviceName + "] Ping OK")
			}
		}
	}()

	return cancel
}
