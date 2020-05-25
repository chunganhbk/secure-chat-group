package main

import (
	"github.com/chunganhbk/chat-golang/database"
	"log"
	"net/http"
	"os"
)

func main(){

	//Setup database
	//init router
	port := os.Getenv("PORT")
	router := NewRouter()


	database.DB = database.SetupDB()
	defer database.DB.Close()

	//create controller server
	log.Fatal(http.ListenAndServe(":"+port, router))
}


