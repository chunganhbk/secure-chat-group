package main

import (
	"github.com/chunganhbk/chat-golang/database"
	"log"
	"net/http"
	"os"
)

func main(){

	//Setup database
	database.SetupDB()
	//init router
	port := os.Getenv("PORT")
	router := NewRouter()





	//create controller server
	log.Fatal(http.ListenAndServe(":"+port, router))
}


