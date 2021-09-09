package middleware

import "github.com/gin-gonic/gin"

func Register(e *gin.Engine) {
	e.Use(requestId())
	e.Use(cors())
}

