package app

import (
	router "docker-api/app/docker/router"
	"docker-api/app/user"
	"docker-api/common/middleware"
	_ "docker-api/docs"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func Register(e *gin.Engine) {
	// Example ping request.
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "result": gin.H{"time": time.Now().Unix(), "requestId": requestid.Get(c)}})
	})
	e.POST("/auth", user.Login)
	api := e.Group("/api/v1")
	api.Use(middleware.Jwt())
	router.RegisterDockerGroup(api)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
