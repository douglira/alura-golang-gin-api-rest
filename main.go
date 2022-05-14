package main

import (
	"github.com/douglira/alura-golang-gin-api-rest/database"
	"github.com/douglira/alura-golang-gin-api-rest/routes"

	"github.com/douglira/alura-golang-gin-api-rest/validators"
)

func main() {
	validators.LoadValidators()
	database.ConnectDatabase()
	r := routes.GetRouter()
	r.Run(":5000")
}
