package middleware

import (
	j "docker-api/common/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "缺少token"})
			c.Abort()
			return
		}
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("无效token| %v", err)})
			c.Abort()
			return
		}
		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
