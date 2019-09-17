package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"
)

var DB db.Persist

type User = models.User

func init() {
	DB = db.PostGresSQL{}
	var err error
	DB, err = DB.Connect()
	if err != nil {
		log.Panic(err)
	}
}

func CreateUser(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := u.Valid(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = DB.SaveUser(u)
	if err != nil {
		log.Println("DB is nil", DB)
		return
	}

	c.JSON(http.StatusOK, u)
}

func GetUser(c *gin.Context) {
	us, _ := DB.GetUser()
	c.JSON(http.StatusOK, us)
}
