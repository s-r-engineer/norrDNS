package main

import (
	"io"
	"net/http"

	libraryErrors "github.com/s-r-engineer/library/errors"
)

func parseBody(resp *http.Response) (b []byte) {
	b, err := io.ReadAll(resp.Body)
	libraryErrors.Errorer(err)
	return
}

func compareSlices(s1, s2 []string) bool {
	if len(s1) == len(s2) {
		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}
	}
	return false
}
