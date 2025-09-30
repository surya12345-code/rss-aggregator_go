package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth error")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth error")
	}
	return vals[1], nil

}
