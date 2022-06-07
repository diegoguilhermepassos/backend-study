package main

import (
	"github.com/hyperguri/webapi-with-go/database"
	"github.com/hyperguri/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
