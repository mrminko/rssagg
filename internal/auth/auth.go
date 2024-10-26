package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("api key not included")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("api key wrong format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("api key wrong structure")
	}
	return vals[1], nil
}
