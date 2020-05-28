package main

import (
	"github.com/chunganhbk/chat-golang/database"
	"github.com/chunganhbk/chat-golang/handlers"
	"github.com/chunganhbk/chat-golang/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	HubGlob *handlers.Hub
)

func main() {
	err := godotenv.Load()
	var log = logrus.New()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Setup database
	var dataStore = database.NewDataStore(log)
	//init router
	r := gin.Default()
	r.Use(middleware.VerifySiteMiddleware(dataStore))
	HubGlob = handlers.NewHub()
	go HubGlob.Run()

	r.GET("/ws", handlers.WebSocketUpgrade)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := r.Group("/api/v1").Use(middleware.CORSMiddleware())
	{
		api.POST("/channel", handlers.CreateChannelRoute)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)

	//create handlers server
	//log.Fatal(http.ListenAndServe(":"+port, router))
}
