package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func requestId() gin.HandlerFunc{
	return requestid.New(requestid.Config{
		//Generator: func() string {
		//	return "test"
		//},
	})
}
