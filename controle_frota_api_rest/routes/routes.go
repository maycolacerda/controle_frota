package routes

import (
	"controle_frota_gin/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

// It creates a new Gin router, loads the HTML templates, defines the routes, and starts the server
func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*.html")

	// A route to the index page.
	r.GET("/", controllers.Index)
	r.GET("/home", controllers.RenderIndex)

	//usuários
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id_user", controllers.GetUser)
	r.POST("/users/newuser", controllers.NewUser)

	//veiculos
	r.GET("/veiculos/:id_veiculo", controllers.GetVeiculo)
	r.GET("/veiculos", controllers.Veiculos)
	r.POST("/novoveiculo", controllers.NovoVeiculo)
	r.DELETE("/deletarveiculo/:id_veiculo", controllers.DeletarVeiculo)
	r.PATCH("/atualizarveiculo/:id_veiculo", controllers.AtualizarVeiculo)

	//motoristas
	r.GET("/motoristas", controllers.Motoristas)
	r.GET("/motoristas/:id_motorista", controllers.GetMotorista)
	r.POST("/novomotorista", controllers.NovoMotorista)
	r.PATCH("/atualizarmotorista/:id_motorista", controllers.AtualizarMotorista)

	//servicos
	r.GET("/servicos", controllers.Servicos)
	r.GET("/servicos/:id_servico", controllers.GetServico)
	r.POST("/novoservico", controllers.NovoServico)
	r.PATCH("/atualizarservico/:id_servico", controllers.AtualizarServico)

	//Manutencoes
	r.GET("/manutencoes", controllers.Manutencoes)
	r.GET("/manutencoes/:id_manutencao", controllers.GetManutencao)
	r.POST("/novamanutencao", controllers.NovaManutencao)
	r.PATCH("/atualizarmanutencao/:id_manutencao", controllers.AtualizarManutencao)
	//404
	r.NoRoute(controllers.NotFound)

	log.Fatal(http.ListenAndServe("localhost:8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) //libera o consumo da api através de qualquer porta]

}
