package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/miekg/dns"

	libraryErrors "github.com/s-r-engineer/library/errors"
	libraryLogging "github.com/s-r-engineer/library/logging"
)

var FQDNBase []string

func initParams() {
	domain := os.Getenv("NORRDNS_REQUEST_DOMAIN")
	if !regexp.MustCompile(`(?i)^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,63}\.?$`).MatchString(domain) {
		panic(os.Getenv("NORRDNS_REQUEST_DOMAIN"))
	}
	FQDNBase = strings.Split(domain, ".")
}

func main() {
	initClient()
	getCountries()
	//initDatabase() // off by default
	initParams()
	libraryLogging.Info("init done")
	dns.HandleFunc(".", handleDNSRequest)
	server := &dns.Server{Addr: ":5333", Net: "udp"}
	libraryErrors.Errorer(server.ListenAndServe())
}
