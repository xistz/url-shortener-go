package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Message struct
type Message struct {
	Message string `json:"message"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{"hello"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	http.HandleFunc("/", rootHandler)

	log.Println("Listenining on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
