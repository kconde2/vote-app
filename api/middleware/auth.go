package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
		LoginResponse: LoginResponse,
	})
}

// Callback function that should perform the authorization of the authenticated user. Called
// only after an authentication success. Must return true on success, false on failure.
// Optional, default to success.
func authorizator (data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok && v.Email == "admin" {
		return true
	}

	return false
}

// Callback function that should perform the authentication of the user based on login info.
// Must return user data as user identifier, it will be stored in Claim Array. Required.
// Check error (e) to determine the appropriate error message.
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

		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return nil,jwt.ErrFailedAuthentication
			// return 
		} else {
			return &models.User{
				Email:  email,
				UUID:  user.UUID,
			}, nil
		}
}

func identityHandler (c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		Email: claims[identityKey].(string),
	}
}

// Callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func payloadFunc (data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
}

// User can define own LoginResponse func.
func LoginResponse (c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, gin.H{
			"jwt":  token,
	})
}
