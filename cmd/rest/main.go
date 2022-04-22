package main

import (
	"core-auth/internal/app/server"
)

// import (
// 	"context"
// 	"log"
// 	// "github.com/Nerzal/gocloak/v11"
// )

func main() {
	server.NewServer()
	// var clientID = "backoffice-auth"
	// var clientSecret = "NF1mBPTsxOGBW4IJBYNlXDfd8SXWGAnj"
	// var realm = "master"

	// // client := gocloak.NewClient("http://localhost:8080")
	// client := gocloak.NewClient("http://localhost:8080", gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
	// ctx := context.Background()
	// token, err := client.LoginClient(ctx, clientID, clientSecret, realm)
	// if err != nil {
	// 	panic("Login failed:" + err.Error())
	// }
	// log.Println(token)

}
