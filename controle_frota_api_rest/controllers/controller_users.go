package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {

	var user models.Users

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}
	err = models.ValidacaoUsers(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else {
		database.DB.Create(&user)
		c.JSON(http.StatusOK, user)
	}

}
