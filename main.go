package main

import (
	"os"
	"strings"

	"github.com/miekg/dns"

	libraryErrors "github.com/s-r-engineer/library/errors"
	libraryLogging "github.com/s-r-engineer/library/logging"
)

var FQDNBase []string

func initParams() {
	domain := os.Getenv("NORRDNS_REQUEST_DOMAIN")
	FQDNBase = strings.Split(domain, ".")
}

func main() {
	initClient()
	getCountries()
	//initDatabase() // oof by default
	initParams()
	libraryLogging.Info("init done")
	dns.HandleFunc(".", handleDNSRequest)
	server := &dns.Server{Addr: ":5333", Net: "udp"}
	libraryErrors.Errorer(server.ListenAndServe())
}
