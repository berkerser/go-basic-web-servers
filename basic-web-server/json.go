package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func helloJSON(mux *http.ServeMux) {
	mux.HandleFunc("/json", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		// Reads entire json file and buffers it
		f, err := ioutil.ReadFile("json/hello.json")
		if err != nil {
			log.Println(err)
			return
		}
		rw.Write(f)
	})
}
