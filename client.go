package main

import (
	"github.com/chunganhbk/chat-golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

type Claims struct {
	User models.UserSchema
	jwt.StandardClaims
}
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	user *Claims
	send chan []byte
}
