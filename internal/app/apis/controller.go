package apis

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller type definition
type Controller struct {
	keycloak *keycloak
}

// NewController returns new controller instance.
func NewController(e *gin.Engine, keycloak *keycloak) Controller {
	ctrl := Controller{keycloak}

	e.POST("/sign-in", ctrl.SignIn)

	return ctrl
}

func (ctrl *Controller) SignIn(c *gin.Context) {
	client := ctrl.keycloak.gocloak
	ctx := context.Background()
	token, err := client.LoginClient(ctx, ctrl.keycloak.clientId, ctrl.keycloak.clientSecret, ctrl.keycloak.realm)
	if err != nil {
		panic("Login failed:" + err.Error())
	}
	c.JSON(http.StatusOK, token)
}
