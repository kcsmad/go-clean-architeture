package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Initializing API Server...")

	router := mux.NewRouter()

	router.HandleFunc("/", )

	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal(err)
	}
}