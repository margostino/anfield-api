package auth

import (
	"net/http"
	"os"
	"strings"
)

var validApiTokens = loadApiTokens()

func loadApiTokens() map[string]bool {
	apiTokensMap := make(map[string]bool)
	apiTokensEnvVar := os.Getenv("API_TOKENS")

	if apiTokensEnvVar == "" {
		return apiTokensMap
	}

	apiTokens := strings.Split(apiTokensEnvVar, ",")

	for _, token := range apiTokens {
		apiTokensMap[token] = true // TODO: implement proper revoke
	}

	return apiTokensMap
}

func IsAuthorized(r *http.Request) bool {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return false
	}
	requestApiToken := authorizationHeader[7:]
	isValidApiToken := validApiTokens[requestApiToken]
	return isValidApiToken
}
