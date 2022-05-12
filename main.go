package main

import (
	"github.com/douglira/alura-golang-gin-api-rest/database"
	"github.com/douglira/alura-golang-gin-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	r := routes.GetRouter()
	r.Run(":5000")
}
