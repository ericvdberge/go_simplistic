package services

import (
	"context"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

func LoginUserWithUserNameAndPassword(username string, password string) (*gocloak.JWT, error) {
	clientId := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_REALM")
	url := os.Getenv("KEYCLOAK_URL")

	client := gocloak.NewClient(url)
	ctx := context.Background()
	return client.Login(ctx, clientId, clientSecret, realm, username, password)
}

func LoginUserWithJWT(jwt string) bool {
	clientId := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_REALM")
	url := os.Getenv("KEYCLOAK_URL")

	client := gocloak.NewClient(url)
	ctx := context.Background()
	tokenResult, error := client.RetrospectToken(ctx, jwt, clientId, clientSecret, realm)
	return *tokenResult.Active && error == nil
}
