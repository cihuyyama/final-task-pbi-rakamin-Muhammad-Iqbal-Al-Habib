package main

import (
	"task-5-pbi-btpns-habib/database"
	"task-5-pbi-btpns-habib/initializers"
	"task-5-pbi-btpns-habib/router"
)

func main() {
	initializers.LoadEnvVariables()

	database.Connect()

	database.Migrate()

	r := router.SetupRouter()

	r.Run()
}
