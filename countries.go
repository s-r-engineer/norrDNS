package main

import (
	"encoding/json"
	"net/http"
	"strings"

	libraryErrors "github.com/s-r-engineer/library/errors"
	libraryLogging "github.com/s-r-engineer/library/logging"
	librarySync "github.com/s-r-engineer/library/sync"
)

var countriesLive map[string]int
var lock, unlock = librarySync.GetMutex()

func init() {
	countriesLive = map[string]int{}
}

func countrieExist(s string) int {
	lock()
	defer unlock()
	v, ok := countriesLive[s]
	if !ok {
		return -1
	}
	return v
}

func getCountries() {
	req, err := http.NewRequest("GET", "https://api.nordvpn.com/v1/countries", nil)
	libraryLogging.Dumper(err)
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:132.0) Gecko/20100101 Firefox/132.0")
	resp, err := client.Do(req)
	libraryErrors.Errorer(err)
	defer resp.Body.Close()
	data := parseBody(resp)
	countries := countries{}
	libraryErrors.Errorer(json.Unmarshal(data, &countries))
	lock()
	defer unlock()
	for _, c := range countries {
		countriesLive[strings.ToLower(c.Code)] = c.ID
	}
}

type countries []struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
