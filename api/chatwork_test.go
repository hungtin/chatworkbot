package api_test

import (
	"testing"

	"github.com/hungtin/chatworkbot/api"
)

func TestPostMessage(t *testing.T) {
	cw := api.NewChatworkClient(api.ChatworkToken)
	roomID := 125747327
	err := cw.PostMessage(roomID, "Message sent from test")
	if err != nil {
		t.Error(err)
	}
}
