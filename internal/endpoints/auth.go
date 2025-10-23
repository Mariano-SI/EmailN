package endpoints

import (
	"context"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		keycloakBaseURL := os.Getenv("KEYCLOAK_BASE_URL")
		keycloakRealm := os.Getenv("KEYCLOAK_REALM")
		keycloakClientID := os.Getenv("KEYCLOAK_CLIENT_ID")
		providerURL := keycloakBaseURL + "/realms/" + keycloakRealm

		provider, err := oidc.NewProvider(r.Context(), providerURL)

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "error to connect to provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: keycloakClientID})

		_, err = verifier.Verify(r.Context(), tokenString)

		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
