package apis

import (
	"log"
	"os"

	"github.com/Nerzal/gocloak/v11"
	"github.com/joho/godotenv"
)

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func getKeycloak() *keycloak {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8080", gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms")),
		clientId:     os.Getenv("clientId"),
		clientSecret: os.Getenv("clientSecret"),
		realm:        os.Getenv("realm"),
	}
}
