package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Motoristas(c *gin.Context) {

	var motoristas []models.Motorista
	database.DB.Find(&motoristas)
	c.JSON(200, motoristas)

}

func GetMotorista(c *gin.Context) {

	id := c.Params.ByName("id_motorista")
	var motorista models.Motorista
	database.DB.First(&motorista, id)
	c.JSON(200, motorista)

}

func NovoMotorista(c *gin.Context) {

	var motorista models.Motorista
	if err := c.ShouldBindJSON(&motorista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		if err := models.ValidacaoMotorista(&motorista); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Create(&motorista)
			c.JSON(http.StatusOK, motorista)
		}
	}
}

func AtualizarMotorista(c *gin.Context) {

	var motorista models.Motorista
	id := c.Params.ByName("id_motorista")
	database.DB.First(&motorista, id)
	if err := c.ShouldBindJSON(&motorista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		if err := models.ValidacaoMotorista(&motorista); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Save(&motorista)
			c.JSON(http.StatusOK, motorista)
		}
	}
}
