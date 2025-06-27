package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
