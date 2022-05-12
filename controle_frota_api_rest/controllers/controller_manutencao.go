package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

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

func NovaManutencao(c *gin.Context) {

	var manutencao models.Manutencao
	if err := c.ShouldBindJSON(&manutencao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		if err := models.ValidacaoManutencao(&manutencao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Create(&manutencao)
			c.JSON(http.StatusOK, manutencao)
		}
	}
}

func UpdateManutencao(c *gin.Context) {

	var manutencao models.Manutencao
	id := c.Params.ByName("id_manutencao")
	database.DB.First(&manutencao, id)
	if err := c.ShouldBindJSON(&manutencao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		if err := models.ValidacaoManutencao(&manutencao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Save(&manutencao)
			c.JSON(http.StatusOK, manutencao)
		}
	}
}
