package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	addr          = ":443"
	httpAddr      = ":80"
	key           = "key.pem"
	certification = "cert.pem"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Printing HTTPS.")
	fmt.Println()
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()

	helloJSON(mux)
	// TODO need to handle root better
	mux.HandleFunc("/", rootHandle)

	// To create http server
	// go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), http.HandlerFunc(redirectTLS)))
	go redirectToHTTPS()

	// To create https server
	log.Fatal(http.ListenAndServeTLS(addr, certification, key, mux))
}
