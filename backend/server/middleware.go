package server

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func (s Server) loggedInUser(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenString := headerParts[1]

		token, err := jwt.Parse(tokenString, s.jwks.Keyfunc)
		if err != nil {
			ErrorResponse(w, http.StatusUnauthorized, "Could not parse token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ErrorResponse(w, http.StatusUnauthorized, "claims nahi mile")
			return
		}

		audiences, ok := coerceList[string](claims["aud"])
		if !ok {
			ErrorResponse(w, http.StatusUnauthorized, "audiences incorrect")
			return
		}

		if !slices.Contains(audiences, s.audience) {
			ErrorResponse(w, http.StatusUnauthorized, "Ye wasseypur hai...")
			return
		}

		userId, ok := claims["sub"].(string)
		if !ok {
			ErrorResponse(w, http.StatusUnauthorized, "invalid user id (sub)")
			return
		}

		r.Header.Set("X-USER-ID", userId)

		handler(w, r)
	}
}

func coerceList[T any](list any) ([]T, bool) {
	tList, ok := list.([]T)
	if ok {
		return tList, true
	}

	anyList, ok := list.([]any)
	if !ok {
		return nil, false
	}

	tList = make([]T, len(anyList))
	for i, a := range anyList {
		t, ok := a.(T)
		if !ok {
			return nil, false
		}

		tList[i] = t
	}
	return tList, true
}

func (s Server) logRequest(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request on %s\n", r.Method, r.URL)
		handler(w, r)
	}
}
