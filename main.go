package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Chatwork Bot"))
	})

	fmt.Println("Server start to listen")
	http.ListenAndServe(":80", router)
}
