package router

import (
	"docker-api/app/docker/apis"
	"github.com/gin-gonic/gin"
)

func RegisterDockerGroup(g *gin.RouterGroup) {
	app := g.Group("/app")
	{
		app.GET("/dockers", apis.ListContainer)
		app.POST("/docker/:containerId/stop", apis.StopContainer)
		app.POST("/docker/:containerId/restart", apis.RestartContainer)
	}
}
