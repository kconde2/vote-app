package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	"time"
	// "log"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/kconde2/vote-app/api/models"
	"github.com/kconde2/vote-app/api/db"
)

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "email"

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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
		LoginResponse: LoginResponse,
	})
}

// Callback function that should perform the authorization of the authenticated user. 
func authorizator (data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok && v.Email == "admin" {
		return true
	}

	return false
}

// Callback function that should perform the authentication of the user based on login info.
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
		db.Where("email = ?", email).First(&user)
		
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return nil,jwt.ErrFailedAuthentication
		} 
			return &models.User{
				Email:  email,
				UUID:  user.UUID,
			}, nil
}

func identityHandler (c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		Email: claims[identityKey].(string),
	}
}

// Callback function that will be called during login to define jwt's payload
func payloadFunc (data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"uuid": v.UUID,
					"accessLevel": v.AccessLevel,
				}
			}
			return jwt.MapClaims{}
}

// Defines what is retuned when you call the LoginHandler method
func LoginResponse (c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, gin.H{
			"jwt":  token,
	})
}
