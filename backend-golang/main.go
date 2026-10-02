package main

import (
	"backend-golang/config"
	"backend-golang/database"
	"backend-golang/routes"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	router := routes.SetupRouter()

	router.Run(":3000")
}