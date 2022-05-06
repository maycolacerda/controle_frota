package routes

import (
	"controle_frota_gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/", controllers.Index)
	r.GET("/veiculos/:id_veiculo", controllers.GetVeiculo)
	r.GET("/veiculos", controllers.Veiculos)
	r.POST("/novoveiculo", controllers.NovoVeiculo)
	r.DELETE("/deletarveiculo/:id_veiculo", controllers.DeletarVeiculo)
	r.PATCH("/atualizarveiculo/:id_veiculo", controllers.AtualizarVeiculo)
	r.Run("localhost:8000")

}
