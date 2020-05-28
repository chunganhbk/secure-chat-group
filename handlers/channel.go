package handlers

import (
	"context"
	"fmt"
	"github.com/chunganhbk/chat-golang/models"
	"github.com/chunganhbk/chat-golang/typing"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateType struct {
	Name        string
	PrivateKeys map[string]string
	UserMap     map[string]string
}

func CreateChannelRoute(c *gin.Context) {
	var form CreateType
	c.BindJSON(&form)

	if form.Name == "" {
		c.JSON(400, gin.H{"message": "You must send a name of the channel"})
		return
	}
	var IsSecure = true
	if len(form.UserMap) == 0 {
		c.JSON(400, gin.H{"message": "You must send a map of the user_ids and usernames"})
		return
	}
	if len(form.PrivateKeys) == 0 {
		IsSecure = false
	}

	newChannel := models.ChannelSchema{ID: primitive.NewObjectID(), Name: form.Name, PrivateKeys: form.PrivateKeys,
		IsSecure: IsSecure, UserMap: form.UserMap}

	go (func() {
		message := typing.WebsocketMessageType{MessageType: "NEW_CHANNEL", MessageContent: newChannel}
		var HubGlob *Hub
		clients := make([]string, 0)
		clients = append(clients, newChannel.ID.String())
		HubGlob.createMessage <- CreatedMessageStruct{message: &message, clients: &clients}
	})()
	db := c.MustGet("DB").(*mongo.Database)
	var Channels = db.Collection("channels")
	_, err := Channels.InsertOne(context.TODO(), newChannel)
	if err != nil {
		fmt.Println("Fatal error while creating channel:", err)
		c.JSON(500, gin.H{"message": "Failed to create channel", "err": err})
		return
	}

	c.JSON(200, gin.H{"message": "Channel has been created", "data": newChannel})

}
