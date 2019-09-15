package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	now := time.Now()
	u := User{
		FirstName:   "Henri",
		LastName:    "Lepic",
		DateOfBirth: now,
	}

	fmt.Println(u)

	payload, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(payload))

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, u)
	})

	r.POST("/users", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)

		log.Println(u)
		c.JSON(200, nil)
	})
	r.Run(":8080")
}

// User represents user attributes
type User struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}
