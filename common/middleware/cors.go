package middleware

import (
	c "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func cors() gin.HandlerFunc {
	return c.New(c.Config{
		//AllowAllOrigins: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	})
}