package model

type Notification struct {
	NotificationType string  `json:"notificationType"`
	Content          Content `json:"content"`
}

type Content struct {
	Operation   string `json:"operation"`
	ServiceName string `json:"serviceName"`
	Status      string `json:"status"`
	Detail      string `json:"detail"`
}
