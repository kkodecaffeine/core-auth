package apis

import (
	"core-auth/internal/app/apis/usersvc"

	"core-auth/internal/pkg/user"
	userRepo "core-auth/internal/pkg/user/persistence"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App interface {
	Init() error
	RegisterRoute(driver *gin.Engine)
	Clean() error
}

type apiApp struct {
	db *gorm.DB
}

func (ag *apiApp) Init() error {
	ag.db = getDatabase()
	return nil
}

func (ag *apiApp) RegisterRoute(driver *gin.Engine) {
	nu := user.NewUseCase(userRepo.New(ag.db))
	usersvc.NewController(driver, nu)
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

	r.Run(":8080")
	agApp.Clean()
}
