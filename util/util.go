package util

import (
	"log"
	"net/http"
	"strings"
)

// TrimAllspaces trims all spaces front and back of a string
func TrimAllspaces(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

// CheckErr catches error
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckStatusCode checks status code
func CheckStatusCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed:", resp.StatusCode)
	}
}
