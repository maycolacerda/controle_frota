package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Servicos(c *gin.Context) {

	var servicos []models.Servico
	database.DB.Find(&servicos)
	c.JSON(200, servicos)

}

func GetServico(c *gin.Context) {

	id := c.Params.ByName("id_servico")
	var servico models.Servico
	database.DB.First(&servico, id)
	c.JSON(200, servico)

}

//cria um novo servico verificando os campos do json e se motorista e veiculo existem e se o motorista possui carteira para dividir o veiculo
func NovoServico(c *gin.Context) {

	var servico models.Servico
	err := c.ShouldBindJSON(&servico)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoServico(&servico)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	val := ValidaServico(servico.IdMotorista, servico.IdVeiculo)
	if val {
		database.DB.Create(&servico)
		c.JSON(http.StatusOK, servico)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Motorista ou Veiculo não existe ou não possui carteira para o veiculo",
		})
	}

}

func AtualizarServico(c *gin.Context) {

	var servico models.Servico
	id := c.Params.ByName("id_servico")
	database.DB.First(&servico, id)
	err := c.ShouldBindJSON(&servico)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoServico(&servico)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	val := ValidaServico(servico.IdMotorista, servico.IdVeiculo)
	if val {
		database.DB.Save(&servico)
		c.JSON(http.StatusOK, servico)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Motorista ou Veiculo não existe ou não possui carteira para o veiculo",
		})
	}

}
