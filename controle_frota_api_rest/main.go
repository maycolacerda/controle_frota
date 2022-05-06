package main

import (
	"controle_frota_gin/database"
	"controle_frota_gin/routes"
)

func main() {

	database.Connect()
	routes.HandleRequests()

}
