package handlers

import (
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello handler...")
	w.Write([]byte("hello..."))
}
