package main

import (
	"github.com/gin-gonic/gin"
	"github.com/arhammusheer/performance-testing/golang/config"
)

conf := config.NewConfig()

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.Run()
}
