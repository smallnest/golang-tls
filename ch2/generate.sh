#!/bin/bash

 go run "$(go env GOROOT)/src/crypto/tls/generate_cert.go" --host localhost