package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println(addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
