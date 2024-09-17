package main

import (
	c "context"
	"fmt"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Print("Hello")

		var clientId = r.URL.Query().Get("clientId")
		clientProfile, ok := database[clientId]
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		//fmt.Print("HelloAgain")

		token := r.Header.Get("Authorization")

		fmt.Printf("Token acquired: %s!!\n", token)

		if !isValidToken(clientProfile, token) {
			fmt.Println(" Invalid Token")

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
func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		fmt.Println(strings.TrimPrefix(token, "Bearer "))
		fmt.Println(clientProfile.Token)
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}
