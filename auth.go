package main

import (
	c "context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var clientId = r.URL.Query().Get("clientId")
		clientProfile, ok := database[clientId]
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		token := r.Header.Get("Authorization")
		if !isValidToken(clientProfile, token){
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Store the client profile. this is to avoid calling the db for the same data multiple times
		ctx := c.WithValue(r.Context(), "clientProfile", clientProfile)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

// Note: Bearer token is a token which is passed through the header under the authorization key
func isValidToken(clientProfile ClientProfile, token string) bool{
	if strings.HasPrefix(token, "Bearer "){
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}