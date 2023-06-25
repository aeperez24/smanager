package server

import (
	"smanager/config/db"
	"smanager/config/repository"
	"smanager/internal/user"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.New()
	db := db.DbSqliteConnection()
	userRepository := &repository.GenericGormRepository[user.User]{db}
	service := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(service)
	r.POST("/user/", userHandler.Create)
	return r
}
