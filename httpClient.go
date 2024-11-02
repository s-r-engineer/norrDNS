package main

import (
	"crypto/tls"
	"net/http"
)

var client *http.Client

func initClient() {
	tlsConfig := &tls.Config{}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client = &http.Client{
		Transport: transport,
	}
}
