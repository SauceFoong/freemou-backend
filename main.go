package main

import (
	"freemou-backend/configs"
	"freemou-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	
	//connect with db
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	router.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	router.Run("localhost:6000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}