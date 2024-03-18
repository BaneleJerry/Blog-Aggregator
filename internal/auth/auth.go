package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthIncluded = errors.New("no authorization header included")

// Get apiKey from Auth header
func GetApiKey(header *http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthIncluded
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) < 2 || authSplit[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return authSplit[1], nil
}
