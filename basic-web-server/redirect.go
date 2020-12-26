package main

import (
	"log"
	"net/http"
	"os"
)

func redirectToHTTPS() {
	srv := http.Server{
		Addr: httpAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := r.URL
			u.Scheme = "https"

			// If there is request from local.
			if u.Host == "" {
				u.Host = os.Getenv("IP")
			}

			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}

	log.Fatal(srv.ListenAndServe())
}
