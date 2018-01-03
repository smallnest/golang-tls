package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tls"))
	})

	var srv http.Server
	srv.Addr = "localhost:8864"
	http2.ConfigureServer(&srv, &http2.Server{})

	err := srv.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
