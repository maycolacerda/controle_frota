package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// It gets all the motoristas from the database and returns them as JSON
func Motoristas(c *gin.Context) {

	var motoristas []models.Motorista
	database.DB.Find(&motoristas)
	c.JSON(200, motoristas)

}

// It gets the id from the URL, then it gets the motorista from the database, and then it returns the
// motorista as JSON
func GetMotorista(c *gin.Context) {

	id := c.Params.ByName("id_motorista")
	var motorista models.Motorista
	database.DB.First(&motorista, id)
	c.JSON(200, motorista)

}

// It receives a JSON object, binds it to a struct, validates the struct and then inserts it into the
// database.
func NovoMotorista(c *gin.Context) {

	var motorista models.Motorista
	err := c.ShouldBindJSON(&motorista)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoMotorista(&motorista)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		database.DB.Create(&motorista)
		c.JSON(http.StatusOK, motorista)
	}

}

// It receives a JSON object, binds it to a struct, validates the struct, and then saves it to the
// database
func AtualizarMotorista(c *gin.Context) {

	var motorista models.Motorista
	id := c.Params.ByName("id_motorista")
	database.DB.First(&motorista, id)
	if err := c.ShouldBindJSON(&motorista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err := models.ValidacaoMotorista(&motorista)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		database.DB.Save(&motorista)
		c.JSON(http.StatusOK, motorista)
	}

}
