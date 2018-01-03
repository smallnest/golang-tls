package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	transport := &http2.Transport{
		AllowHTTP: true,
		// TLSClientConfig: &tls.Config{
		//  InsecureSkipVerify: true,
		// },
		DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(netw, addr)
		},
	}

	client := http.Client{Transport: transport}
	resp, err := client.Get("https://localhost:8864/hello")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("is HTTP2: %v (%s)\n\n", resp.ProtoAtLeast(2, 0), resp.Proto)

	io.Copy(os.Stdout, resp.Body)
}
