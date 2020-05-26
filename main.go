package main

import (
	"github.com/chunganhbk/chat-golang/database"
	"github.com/chunganhbk/chat-golang/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	//Setup database
	database.SetupDB()
	//init router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := r.Group("/api/v1")
	{

	}

	port := os.Getenv("PORT")
	r.Run(":" + port)

	//create controller server
	//log.Fatal(http.ListenAndServe(":"+port, router))
}
