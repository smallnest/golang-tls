package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tls"))
	})

	tlsConfig := &tls.Config{
		// NoClientCert
		// RequestClientCert
		// RequireAnyClientCert
		// VerifyClientCertIfGiven
		// RequireAndVerifyClientCert
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
