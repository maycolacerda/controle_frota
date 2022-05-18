package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Manutencoes(c *gin.Context) {

	var manutencoes []models.Manutencao
	database.DB.Find(&manutencoes)
	c.JSON(200, manutencoes)

}

func GetManutencao(c *gin.Context) {
	id := c.Params.ByName("id_manutencao")
	var manutencao models.Manutencao
	database.DB.First(&manutencao, id)
	c.JSON(200, manutencao)
}

//cria uma nova manutençao verificando se veiculo existe

func NovaManutencao(c *gin.Context) {

	var manutencao models.Manutencao
	idveiculo := c.Params.ByName("id_veiculo")
	err := c.ShouldBindJSON(&manutencao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoManutencao(&manutencao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	temp, err := strconv.ParseUint(idveiculo, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	if ValidaVeiculo(uint(temp)) {
		database.DB.Create(&manutencao)
		c.JSON(http.StatusOK, manutencao)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Veiculo não encontrado",
		})
	}

}

//Atualiza manutenção verificando se veiculo existe

func AtualizarManutencao(c *gin.Context) {

	var manutencao models.Manutencao
	id := c.Params.ByName("id_manutencao")
	idveiculo := c.Params.ByName("id_veiculo")
	database.DB.First(&manutencao, id)

	err := c.ShouldBindJSON(&manutencao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoManutencao(&manutencao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	temp, err := strconv.ParseUint(idveiculo, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	if ValidaVeiculo(uint(temp)) {
		database.DB.Save(&manutencao)
		c.JSON(http.StatusOK, manutencao)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": "Veiculo não encontrado",
		})
	}

}
