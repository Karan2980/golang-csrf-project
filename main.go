package main

import "log"

var host = "localhost"
var port = "9000"

func main() {
	db.InitDB()

	jwtErr := myJwt.InitJET()
	if jwtErr != nil {
		log.Println("Error initializing the JWT!")
	}
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error initializing the server!")
		log.Fatal(serverErr)
	}

}