package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Chatwork Bot"))
	})

	fmt.Println("Server start to listen on ", port)
	http.ListenAndServe(":"+port, router)
}
