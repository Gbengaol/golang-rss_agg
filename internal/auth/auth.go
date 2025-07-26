package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from a request's header
// Example: Authorization: ApiKey {api_key_value_is_here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("authorization key not provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("authorization key is malformed")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("authorization key is malformed")
	}

	return vals[1], nil
}
