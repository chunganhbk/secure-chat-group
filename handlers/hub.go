package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chunganhbk/chat-golang/typing"
	"github.com/mitchellh/mapstructure"
)

type Hub struct {
	clients       map[*Client]bool
	clientMap     map[string]*Client // user [user id] to Client
	receive       chan []byte
	createMessage chan CreatedMessageStruct
	register      chan *Client
	unregister    chan *Client
}

type CreatedMessageStruct struct {
	message *typing.WebsocketMessageType
	clients *[]string // user ids
}

type MESSAGE struct {
	Type    string
	Content map[string]interface{}
}

type IsTyping struct {
	ChannelID         string
	Users             []string
	WhoTypingID       string
	WhoTypingUsername string
}

type OFFER struct {
	Users       []string
	WhoID       string
	WhoUsername string
	CallType    string
	ChannelId   string
	ChannelName string
	SignalData  interface{}
}

func NewHub() *Hub {
	return &Hub{
		createMessage: make(chan CreatedMessageStruct),
		receive:       make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		clientMap:     make(map[string]*Client),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.clientMap[client.user.User.ID.Hex()] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				delete(h.clientMap, client.user.User.ID.Hex())
				fmt.Println(client.send)
				close(client.send)
			}
		case message := <-h.receive:
			handleReceive(h, message)
		case message := <-h.createMessage:
			bytes, err := json.Marshal(message.message)

			if err != nil {
				fmt.Println("Bad", err)
				continue
			}

			for _, client := range *message.clients {
				var connection = h.clientMap[client]
				if connection != nil {
					select {
					case connection.send <- bytes:
					default:
						close(connection.send)
						delete(h.clients, connection)
						delete(h.clientMap, client)
					}
				}
			}
		}
	}
}

func handleReceive(h *Hub, message []byte) {
	var parsed MESSAGE
	json.Unmarshal(message, &parsed)

	switch parsed.Type {
	case "IS_TYPING":
		var content IsTyping

		mapstructure.Decode(parsed.Content, &content)

		message := typing.WebsocketMessageType{MessageType: "IS_TYPING", MessageContent: &content}

		clientsCopy := content.Users[:]

		content.Users = []string{}

		go (func() {
			h.createMessage <- CreatedMessageStruct{message: &message, clients: &clientsCopy}
		})()
	case "OFFER":
		var content OFFER

		mapstructure.Decode(parsed.Content, &content)

		message := typing.WebsocketMessageType{MessageType: "OFFER", MessageContent: &content}

		clientsCopy := content.Users[:]

		content.Users = []string{}

		go (func() {
			h.createMessage <- CreatedMessageStruct{message: &message, clients: &clientsCopy}
		})()
	case "ANSWER":
		var content OFFER

		mapstructure.Decode(parsed.Content, &content)

		message := typing.WebsocketMessageType{MessageType: "ANSWER", MessageContent: &content}

		clientsCopy := content.Users[:]

		content.Users = []string{}

		go (func() {
			h.createMessage <- CreatedMessageStruct{message: &message, clients: &clientsCopy}
		})()
	}

}
