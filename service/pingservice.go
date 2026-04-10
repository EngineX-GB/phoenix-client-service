package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func StartPinger(serviceName string, serverURL string) (stop func()) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("[INFO] [Service : " + serviceName + "] Pinger stopped")
				return

			case <-ticker.C:
				resp, err := http.Get(serverURL + "/ping")
				if err != nil {
					fmt.Println("[ERROR] [Service : "+serviceName+"] Ping error:", err)
					continue
				}
				resp.Body.Close()
				fmt.Println("[INFO] [Service : " + serviceName + "] Ping OK")
			}
		}
	}()

	return cancel
}
