package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kconde2/vote-app/api/controllers"
	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/middleware"
)

func main() {
	log.Println("Starting server...")

	db.Initialize()

	r := gin.Default()

	// Generates JWT token
	authware, err := middleware.AuthMiddleware()
	if err != nil {
	  log.Fatal("JWT Error:" + err.Error())
	}
	v1 := r.Group("/")
	{
		v1.POST("/login", authware.LoginHandler)
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
	test := authware.GetClaimsFromJWT
	log.Println("claim here",test)

	r.Run(":8080")
}
