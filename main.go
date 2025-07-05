package main

import (
	"log"

	"github.com/Karan2980/golang-csrf-project/db"
	"github.com/Karan2980/golang-csrf-project/server"
	myjwt "github.com/Karan2980/golang-csrf-project/server/middleware/myJwt"
)

var host = "localhost"
var port = "9000"

func main() {
	db.InitDB()

	jwtErr := myjwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT!")
	}
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error initializing the server!")
		log.Fatal(serverErr)
	}

}