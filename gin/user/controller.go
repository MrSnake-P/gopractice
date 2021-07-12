package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello world!"})
}

func Login(c *gin.Context) {
	user, err := GetUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "failed"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": user})
}