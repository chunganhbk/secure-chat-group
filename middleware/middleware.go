package middleware

import (
	"context"
	"github.com/chunganhbk/chat-golang/database"
	"github.com/chunganhbk/chat-golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,"+
			" X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, Secret-Key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func VerifySiteMiddleware(dataStore *database.MongoDataStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := c.Request.Header.Get("Secret-Key")
		if len(secretKey) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Secret Key found"})
			return
		}
		var site models.SiteSchema
		err := dataStore.DBMaster.Collection("sites").FindOne(context.TODO(), bson.M{"secret": secretKey}).Decode(&site)
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not query messages", "err": err})
			return
		}
		db := dataStore.Session.Database(site.Database)
		c.Set("db", db)
		c.Next()
	}
}
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if len(tokenString) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No JWT token found"})
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := VerifyToken(tokenString)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Your token is invalid"})
				return
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request", "err": err})
			return
		}

		//todo: find a better way to convert the claim to string
		c.Set("token", claims)
		c.Next()
	}
}
