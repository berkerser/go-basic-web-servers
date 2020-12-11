package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	port          = "PORT"
	key           = "key.pem"
	certification = "cert.pem"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", rootHandle)

	// To gain information about inputs for error debugging
	log.Println("Listening on:")
	log.Printf("PORT  : %s\n", os.Getenv(port))
	log.Printf("ADRESS: /\n")

	// // To create http server
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv(port)), nil))op

	// To create https server
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", os.Getenv(port)), certification, key, nil))
}
