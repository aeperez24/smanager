package server

import (
	"smanager/config/db"
	"smanager/config/repository"
	"smanager/internal/httputils"
	"smanager/internal/login"
	"smanager/internal/managedsecret"
	"smanager/internal/middleware"
	"smanager/internal/middleware/auth"
	"smanager/internal/token"
	"smanager/internal/user"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	//database initialization
	r := gin.New()
	dbConnection := db.DbSqliteConnectionWithFile("localFile")
	db.Migrate(dbConnection)

	//repository initialization
	userRepo := &repository.GenericGormRepository[user.User]{DB: dbConnection}
	managedSecretRepo := &repository.GenericGormRepository[managedsecret.ManagedSecret]{DB: dbConnection}

	//Services initialization
	managedSecretService := managedsecret.NewManagedSercertService(managedSecretRepo)
	userService := user.NewUserService(userRepo)
	tokenService := token.NewTokenService("my internal secret key//TODO MOVE TO ENV VAR")
	loginService := login.NewLoginService(userService, tokenService)

	//Middlewares
	middlewareMaps := map[middleware.MiddlewareType]gin.HandlerFunc{
		middleware.Secured: auth.NewAuthMiddleware(tokenService),
	}
	//handlers configuration
	handlersConfigs := []httputils.HandlerProvider{
		managedsecret.NewHandlerConfigProvider(managedSecretService),
		login.NewLoginHandlerConfigProvider(loginService),
		user.NewLoginHandlerConfigProvider(userService),
	}

	//handlers register

	for _, handlerConfig := range handlersConfigs {
		httputils.RegisterRoutesWithMiddleware(r, handlerConfig.GetHandlers(), middlewareMaps)
	}
	return r
}
