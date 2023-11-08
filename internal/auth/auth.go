package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Shouldbe given in format Authorization: APIKey {apikey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authication found")
	}

	vals := strings.Split(val, " ")
	//check to ensure formatted correctly ie APIKey {apikey}
	if len(vals) != 2 {
		return "", errors.New("malformed key")
	}

	if vals[0] != "APIKey" {
		return "", errors.New("malformed first part of header")
	}

	return vals[1], nil

}
