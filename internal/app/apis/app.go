package apis

import (
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

type App interface {
	Init()
	RegisterRoute(driver *gin.Engine)
	Clean() error
}

type apiApp struct {
	keycloak *keycloak
	// db *gorm.DB
}

func (ag *apiApp) Init() {
	ag.keycloak = getKeycloak()
}

func (ag *apiApp) RegisterRoute(driver *gin.Engine) {
	// nu := user.NewUseCase(userRepo.New(ag.db))
	NewController(driver, ag.keycloak)
}

func (ag *apiApp) Clean() error {
	return nil
}

var agApp *apiApp

// CreateAPIApp returns new core.App implementation.
func CreateAPIApp() {
	r := gin.Default()

	agApp = &apiApp{}
	agApp.Init()
	agApp.RegisterRoute(r)

	r.Run(":3000")
	agApp.Clean()
}
