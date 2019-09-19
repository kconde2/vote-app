package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kconde2/vote-app/api/handlers"
)

var (
	port string
)

func init() {
	// port = os.Getenv("PORT")
	port = "8080"
	if len(port) == 0 {
		log.Panic("no given port")
	}
}

func main() {
	r := gin.Default()
	r.GET("/users", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.Run(":8080")
}
