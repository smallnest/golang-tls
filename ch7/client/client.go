package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// tlsConfig := &tls.Config{
	// 	InsecureSkipVerify: true,
	// }

	// transport := &http2.Transport{TLSClientConfig: tlsConfig}
	// client := &http.Client{Transport: transport}

	// resp, err := client.Get("https://localhost:8864/hello")

	resp, err := http.Get("https://localhost:8864/hello")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("is HTTP2: %v (%s)\n\n", resp.ProtoAtLeast(2, 0), resp.Proto)

	io.Copy(os.Stdout, resp.Body)
}
