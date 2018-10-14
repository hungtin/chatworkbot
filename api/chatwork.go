package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ChatworkBaseURL entry point base to the api
const ChatworkBaseURL = "https://api.chatwork.com/v2"

// A Chatwork struct is a client object
// to interact with the api
type Chatwork struct {
	token string
}

// NewChatworkClient to create a object
// and hold the token
func NewChatworkClient(token string) Chatwork {
	return Chatwork{
		token: token,
	}
}

// prepareReq help to create Request
// method should be GET, POST, PUT, DELETE
// endpoint start WITH first slash
func (cw Chatwork) prepareReq(method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		ChatworkBaseURL+endpoint,
		body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-ChatWorkToken", cw.token)

	return req, nil
}

// GetMe request GET on /me uri
func (cw Chatwork) GetMe() ([]byte, error) {
	client := &http.Client{}
	req, err := cw.prepareReq("GET", "/me", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// PostMessage send message to a room
func (cw Chatwork) PostMessage(roomID int, message string) error {
	client := &http.Client{}

	endpoint := fmt.Sprintf("/rooms/%v/messages", roomID)
	data := url.Values{}
	data.Set("body", message)

	req, err := cw.prepareReq("POST", endpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	// TODO: add return value like message id
	// or anything to clarify that message sent successfully

	return nil
}
