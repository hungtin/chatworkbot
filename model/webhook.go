package model

// A WebhookEvent is a data type
// that contains related data
// about chatwork webhook
type WebhookEvent struct {
	FromAccountID int    `json:"from_account_id"`
	ToAccountID   int    `json:"to_account_id"`
	RoomID        int    `json:"room_id"`
	MessageID     string `json:"message_id"`
	Body          string `json:"body"`
	SendTime      int    `json:"send_time"`
	UpdateTime    int    `json:"update_time"`
}
