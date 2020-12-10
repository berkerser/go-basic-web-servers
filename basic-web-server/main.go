package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	port = "PORT"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", rootHandle)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(port)), nil))
}
