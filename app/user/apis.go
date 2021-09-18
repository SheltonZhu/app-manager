package user

import (
	hash "docker-api/common/bcrypt"
	"docker-api/common/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

var (
	defaultAcc = "docker.admin"
	hashStr    = "$2a$11$XPxyTOksR.LXqqjVEs1vCefPFwGFruNXqZB.qkscfGQz90OTMNePy"
)

func Login(c *gin.Context) {
	user := User{
		c.PostForm("account"),
		c.PostForm("password"),
	}
	if user.Account == "" && user.Password == "" {
		_ = c.BindJSON(&user)
	}
	envPasswd := os.Getenv("DefaultPassword")
	envAcc := os.Getenv("DefaultAccount")
	if envAcc != "" {
		defaultAcc = envPasswd
	}
	if envPasswd != "" {
		hashStr, _ = hash.HashPassword(envPasswd)
	}

	if ok := hash.CheckPasswordHash(user.Password, hashStr); ok && user.Account == defaultAcc {
		token, err := jwt.GenerateToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Generate token failed."})
			log.Panicln(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok", "result": gin.H{"token": token}})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter."})
}
