package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hungtin/chatworkbot/api"

	"github.com/gorilla/mux"
	"github.com/hungtin/chatworkbot/model"
)

func parseWebhookEvent(raw []byte) (*model.WebhookEvent, error) {
	var obj map[string]*json.RawMessage
	err := json.Unmarshal(raw, &obj)
	if err != nil {
		return nil, err
	}

	var eventObj = new(model.WebhookEvent)
	err = json.Unmarshal(*obj["webhook_event"], eventObj)
	if err != nil {
		return nil, err
	}

	return eventObj, nil
}

func chatworkHandlerFunc(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	eventObj, err := parseWebhookEvent(body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(eventObj)
	cw := api.NewChatworkClient("0e7249c219c3afee79a7ada2bfad714c")
	cw.PostMessage(eventObj.RoomID, "What is "+eventObj.MessageID)
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Chatwork Bot"))
	})

	router.HandleFunc("/chatwork", chatworkHandlerFunc)

	fmt.Println("Server start to listen on ", port)
	http.ListenAndServe(":"+port, router)
}
