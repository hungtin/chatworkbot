package api_test

import (
	"reflect"
	"testing"

	"github.com/hungtin/chatworkbot/model"

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

func TestGetMembers(t *testing.T) {
	cw := api.NewChatworkClient(api.ChatworkToken)
	roomID := 125747327
	members, err := cw.GetMembers(roomID)
	if err != nil {
		t.Error(err)
	}

	wantedMembers := []*model.Member{
		&model.Member{AccountID: 2267986, Name: "Trinh Tin[トリン ハン ティン][DEV][ISS担当]"},
		&model.Member{AccountID: 2302663, Name: "楽しいBot"},
	}

	for index, member := range *members {
		if !reflect.DeepEqual(member, wantedMembers[index]) {
			t.Errorf("Want: %v, received: %v", wantedMembers[index], member)
		}
	}
}
