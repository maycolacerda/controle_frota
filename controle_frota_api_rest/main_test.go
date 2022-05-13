package main

import (
	"controle_frota_gin/controllers"
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID uint
)

func SetupRotasTeste(router *gin.Engine) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}
func CriaVeiculoMock() {
	veiculo := models.Veiculo{Identificador: "a5045", Tipo: "Carro", Marca: "Fiat", Modelo: "Uno 2019", Rastreador: true, Placa: "ABC1234", Disponivel: false, Km: true, TipoCarteira: "B"}
	database.DB.Create(&veiculo)
	ID = veiculo.Id
}
func DeletaveiculoMock() {
	var veiculo models.Veiculo
	database.DB.Delete(&veiculo, ID)
}

//teste da home page
func TestStatusHome(t *testing.T) {

	r := SetupRotasTeste(gin.Default())
	r.GET("/", controllers.Index)
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Status deve ser 200")
	mockResposta := `{"mensagem":"Olá mundo","message":"Hello World","status":"ok"}`
	assert.Equal(t, mockResposta, resp.Body.String(), "Mensagem deve ser Olá mundo")
}

//testa a listagem de veiculos
func TestVeiculos(t *testing.T) {
	database.Connect()
	CriaVeiculoMock()
	defer DeletaveiculoMock()
	r := SetupRotasTeste(gin.Default())
	r.GET("/veiculos", controllers.Veiculos)
	req, _ := http.NewRequest("GET", "/veiculos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	//fmt.Println(resp.Body.String())
	assert.Equal(t, 200, resp.Code, "Status deve ser 200")

}
func TestVeiculo(t *testing.T) {

	database.Connect()
	mockResposta := `{"ID":0,"CreatedAt":"2022-05-04T06:41:47Z","UpdatedAt":"2022-05-04T06:41:47Z","DeletedAt":null,"id_veiculo":1,"identificador":"A01","tipo":"Carro","marca":"Ford","modelo":"Fusion","rastreador":true,"placa":"ABC-1234","disponivel":true,"km":true,"tipoCarteira":"B"}`

	r := SetupRotasTeste(gin.Default())
	r.GET("/veiculo/:id_veiculo", controllers.GetVeiculo)
	req, _ := http.NewRequest("GET", "/veiculo/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, mockResposta, resp.Body.String(), "os valores devem ser iguais")

}

func TestManutencao(t *testing.T) {

}
