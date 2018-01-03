package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tls"))
	})

	var srv = &http.Server{}
	srv.Handler = http.DefaultServeMux

	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	http2.ConfigureServer(srv, s2)
	l, _ := net.Listen("tcp", ":8864")
	defer l.Close()
	fmt.Println("Start server...")
	for {
		rwc, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: srv})

	}
}
