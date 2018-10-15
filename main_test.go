package main

import (
	"reflect"
	"testing"

	"github.com/hungtin/chatworkbot/model"
)

func TestParseWebhookEvent(t *testing.T) {
	raw := `{"webhook_setting_id":"2567","webhook_event_type":"mention_to_me","webhook_event_time":1539441501,"webhook_event":{"from_account_id":2267986,"to_account_id":2302663,"room_id":125747327,"message_id":"1102967319673536512","body":"Onegaikoto ga aru desukedo","send_time":1539441501,"update_time":0}}`
	wantedObj := &model.WebhookEvent{
		FromAccountID: 2267986,
		ToAccountID:   2302663,
		RoomID:        125747327,
		MessageID:     "1102967319673536512",
		Body:          "Onegaikoto ga aru desukedo",
		SendTime:      1539441501,
		UpdateTime:    0,
	}
	obj, err := parseWebhookEvent([]byte(raw))
	if err != nil {
		t.Error("Error occurs while run parseWebhookEventTest", err)
	}
	if !reflect.DeepEqual(obj, wantedObj) {
		t.Errorf("Want %v, but received %v", wantedObj, obj)
	}
}

func TestChoseMemberHandler(t *testing.T) {
	eventObj := &model.WebhookEvent{
		FromAccountID: 2267986,
		ToAccountID:   2302663,
		RoomID:        125747327,
		MessageID:     "1102967319673536512",
		Body:          "誰",
		SendTime:      1539441501,
		UpdateTime:    0,
	}

	chooseMemberHandler(eventObj)
}
