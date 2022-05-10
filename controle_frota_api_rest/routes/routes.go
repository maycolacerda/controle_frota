package routes

import (
	"controle_frota_gin/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*.html")

	r.GET("/", controllers.Index)
	r.GET("/home", controllers.RenderIndex)
	r.GET("/veiculos/:id_veiculo", controllers.GetVeiculo)
	r.GET("/veiculos", controllers.Veiculos)
	r.POST("/novoveiculo", controllers.NovoVeiculo)
	r.DELETE("/deletarveiculo/:id_veiculo", controllers.DeletarVeiculo)
	r.PATCH("/atualizarveiculo/:id_veiculo", controllers.AtualizarVeiculo)
	r.NoRoute(controllers.NotFound)
	log.Fatal(http.ListenAndServe("localhost:8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) //libera o consumo da api através de qualquer porta]

}
