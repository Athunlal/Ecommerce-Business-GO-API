package main

import (
	"github.com/athunlal/config"
	"github.com/athunlal/controls"
	"github.com/athunlal/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	config.DBconnect()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/test", controls.UserSignUP)
	r.Run()
}