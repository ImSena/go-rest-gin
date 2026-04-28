package main

import (
	"go-rest-gin/database"
	"go-rest-gin/routes"
)

func main() {
	database.Connection()

	routes.HandleRequests()
}
