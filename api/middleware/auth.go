package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"

	"github.com/kconde2/vote-app/api/models"
	"github.com/kconde2/vote-app/api/db"
)

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "uuid"

// AuthMiddleware : The authmiddlare handling jwt token security
func AuthMiddleware()  (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator: authenticator,
		Authorizator: authorizator,
		Unauthorized: unauthorized,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
}

func unauthorized (c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
}

func authorizator (data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Email == "admin" {
				return true
			}

			return false
}

func authenticator (c *gin.Context) (interface{}, error) {
		var loginVals login
		var user models.User
		var db = db.GetDB()

		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		// request fields
		email := loginVals.Email
		password := loginVals.Password
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		db.Where("email = ?", email).First(&user)
		log.Println("err", err)
		log.Println(user.Password)
		log.Println(hashPassword)

		// // Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return nil,jwt.ErrFailedAuthentication
			// return 
		}

		// Administrator and test user
		if (email == "admin@gmail.com" && password == "admin") || (email == "test@gmail.com" && password == "test") {
			return &models.User{
				Email:  email,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
}

func identityHandler (c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		Email: claims[identityKey].(string),
	}
}

func payloadFunc (data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
}

// handle user login
// func LoginHandler (c *gin.Context){

// 	var creds login
// }
