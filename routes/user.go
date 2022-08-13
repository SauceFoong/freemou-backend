package routes

import (
	"freemou-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here

    router.POST("/user", controllers.CreateUser()) 
    router.GET("/user/:userId", controllers.GetAUser())
    router.PUT("/user/:userId", controllers.UpdateAUser()) 
    router.DELETE("/user/:userId", controllers.DeleteAUser())
    router.GET("/users", controllers.GetAllUsers()) 

}