package main

import (
	"gin/user"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	r := user.Router(g)


	r.Run()
}
