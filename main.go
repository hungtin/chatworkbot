package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func chatworkHandlerFunc(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer r.Body.Close()

	fmt.Println(string(body))
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
