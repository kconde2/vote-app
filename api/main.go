package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kconde2/vote-app/api/controllers"
	"github.com/kconde2/vote-app/api/db"
)

func main() {
	log.Println("Starting server...")

	db.Initialize()

	r := gin.Default()

	v1 := r.Group("/")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
		}
	}

	r.Run(":8080")
}
