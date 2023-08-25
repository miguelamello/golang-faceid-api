package main

import (
	"github.com/miguelamello/golang-faceid-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up Gin in release mode
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Define your routes here
	r.GET("/", routes.GetRoot())
	r.GET("/reference", routes.GetReference())
	r.POST("/frv", routes.PostFRV())

	// Run the server
	r.SetTrustedProxies(nil)
	r.Run("127.0.0.1:8010")
}
