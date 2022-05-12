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

func GetServicos(c *gin.Context) {

	id := c.Params.ByName("id_servico")
	var servico models.Servico
	database.DB.First(&servico, id)
	c.JSON(200, servico)

}

func NovoServico(c *gin.Context) {

	var servico models.Servico
	if err := c.ShouldBindJSON(&servico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		if err := models.ValidacaoServico(&servico); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Create(&servico)
			c.JSON(http.StatusOK, servico)
		}
	}
}
