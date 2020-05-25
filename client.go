package main

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	user *jwt.Claims
	send chan []byte
}