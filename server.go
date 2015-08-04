package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var response = 200

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println(addr)

	go func() {
		<-time.After(2 * time.Minute)
		response = 400
	}()

	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(response)
	io.WriteString(w, "pong")
}
