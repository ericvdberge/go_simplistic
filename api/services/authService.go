package services

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

var (
	clientId     = "public"
	clientSecret = "V9MkdEMTw1QGksOFa95sY7JDvYAtycKY"
	realm        = "myrealm"
)

func LoginUserWithUserNameAndPassword(username string, password string) (*gocloak.JWT, error) {
	client := gocloak.NewClient("http://localhost:8081")
	ctx := context.Background()
	return client.Login(ctx, clientId, clientSecret, realm, username, password)
}

func LoginUserWithJWT(jwt string) bool {
	client := gocloak.NewClient("http://localhost:8081")
	ctx := context.Background()
	tokenResult, error := client.RetrospectToken(ctx, jwt, clientId, clientSecret, realm)
	return *tokenResult.Active == true && error == nil
}
