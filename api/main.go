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
	db.CreateSystemAdmin()

	r := gin.Default()
	route := r.Group("/")

	// Manage login (auth + generate JWT)
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	route.POST("/login", authMiddleware.LoginHandler)

	// Manage protected routes
	route.Use(authMiddleware.MiddlewareFunc())
	{
		users := route.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:uuid", controllers.UpdateUser)
			users.DELETE("/:uuid", controllers.DeleteUser)
		}

		votes := route.Group("/votes")
		{
			votes.GET("/", controllers.GetVotes)
			votes.POST("/", controllers.CreateVote)
			votes.GET("/:uuid", controllers.RetrieveVote)
			votes.PUT("/:uuid", controllers.UpdateVote)
		}
	}

	r.Run(":8080")
}
