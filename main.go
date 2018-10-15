package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/hungtin/chatworkbot/api"

	"github.com/gorilla/mux"
	"github.com/hungtin/chatworkbot/model"
)

func removeTag(name string) string {
	tagReg, err := regexp.Compile("\\[.+?\\]")
	if err != nil {
		log.Fatal(err)
	}
	noTagName := tagReg.ReplaceAllString(name, "")
	return strings.TrimSpace(noTagName)
}

func randomMember(members *[]*model.Member) string {
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(*members))
	return "選ばれた人: " + removeTag((*members)[randomIndex].Name)
}

func chooseMemberHandler(eventObj *model.WebhookEvent) {
	cw := api.NewChatworkClient(api.ChatworkToken)
	var err error
	if strings.Contains(eventObj.Body, "誰") {
		// TODO: Error handler for this
		members, _ := cw.GetMembers(eventObj.RoomID)
		err = cw.PostMessage(eventObj.RoomID, randomMember(members))
	} else {
		err = cw.PostMessage(eventObj.RoomID, "わからん")
	}

	if err != nil {
		log.Println(err)
		cw.PostMessage(eventObj.RoomID, "エラーが発生しました")
	}
}

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
	chooseMemberHandler(eventObj)
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
