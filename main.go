package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/CaioLima42/ascii-conversor-with-web-interface/internal/handlers"
)

func main() {
	route := "localhost:8080"
	r := mux.NewRouter()

	r.HandleFunc("/process", handlers.CreateVideo).Methods("POST")
	r.HandleFunc("/audio", handlers.ExtractAudio).Methods("POST")
	r.HandleFunc("/", handlers.ReadVideo).Methods("GET")

	fmt.Printf("Run in %s\n", route)
	if err := http.ListenAndServe(route, r); err != nil {
		fmt.Println(err.Error())
	}
}
