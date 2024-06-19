package server

import (
	"boilerplate/controllers"
	"boilerplate/db"
	"boilerplate/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(env string) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode)

	databaseLayer := db.NewDatabaseLayer()
	loginService := services.NewLoginService(databaseLayer)
	health := controllers.NewHealthController(loginService)

	router.GET("/status", health.Status)
	router.POST("/login", health.DummyLogin)
	return router
}
