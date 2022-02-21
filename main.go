package main

import (
	"ContactListTask/controllers"
	"ContactListTask/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.InitialDBConnection()

	router.GET("/", controllers.GetContacts)
	router.GET("/:id", controllers.GetContact)
	router.POST("/", controllers.CreateContact)
	router.PUT("/:id", controllers.UpdateContact)
	router.DELETE("/:id", controllers.DeleteContact)

	router.Run(":5000")
}
