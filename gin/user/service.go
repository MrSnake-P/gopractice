package user

import (
	"github.com/gin-gonic/gin"
)


func Router(r *gin.Engine) *gin.Engine {
	r.GET("index", Index)
	user := r.Group("/user")
	{
		user.GET("login", Login)
	}
	return r
}
