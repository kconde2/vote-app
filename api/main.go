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
			users.PUT("/:uuid", controllers.UpdateUser)
			users.DELETE("/:uuid", controllers.DeleteUser)
		}

		votes := v1.Group("/votes")
		{
			votes.GET("/", controllers.GetVotes)
			votes.POST("/", controllers.CreateVote)
			votes.PUT("/:uuid", controllers.UpdateVote)
			votes.DELETE("/:uuid", controllers.DeleteVote)
		}
	}

	r.Run(":8080")
}
