package main

import (
	router "docker-api/app"
	"docker-api/common/middleware"
	"flag"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

//go:generate swag init --parseDependency --parseDepth=6

var (
	httpAddr = flag.String("http", ":4399", "Http listen address.")
)
// @title 容器管理
// @version 1.0
// @description Docker 容器管理
////@termsOfService http://swagger.io/terms/

// @contact.name Shelton Zhu
// @contact.url https://github.com/SheltonZhu
// @contact.email shelton@icloud.com

//// @license.name Apache 2.0
//// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.60.23:4399
// @BasePath /
func main() {
	g := gin.Default()
	middleware.Register(g)
	router.Register(g)
	if err := endless.ListenAndServe(*httpAddr, g); err != nil {
		log.Panicf("error: %s", err)
	}
}
