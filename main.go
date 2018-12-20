package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
