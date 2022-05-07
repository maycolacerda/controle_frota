package main

// @title Controle de Frota
// @version 1.0
// @description sistema para controle de frota.
import (
	"controle_frota_gin/database"
	"controle_frota_gin/routes"
)

func main() {

	database.Connect()
	routes.HandleRequests()

}
