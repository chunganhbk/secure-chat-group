package handlers

import (
	"fmt"
	"github.com/chunganhbk/chat-golang/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func WebSocketUpgrade(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		token = c.GetHeader("Sec-Websocket-Protocol")[5:]
	}

	claims := &Claims{}
	tokenJWT, err := middleware.VerifyToken(token)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Token invalid")
			return
		}

		fmt.Println("Bad request")
		return
	}

	if err = tokenJWT.Valid(); err != nil{
		fmt.Println("Bad token")
		return
	}
	var HubGlob *Hub
	serveWs(HubGlob, c.Writer, c.Request, claims)
}
var protocols = []string{"asd"}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  12800000,
	WriteBufferSize: 12800000,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols:protocols,
}

