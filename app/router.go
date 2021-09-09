package app

import (
	docker "docker-api/app/docker/router"
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
		c.JSON(http.StatusOK, gin.H{"message": "pong", "time": time.Now().Unix(), "requestId": requestid.Get(c)})
	})

	api := e.Group("/api/v1")
	docker.RegisterDockerGroup(api)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
