package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "Hello World",
		"mensagem": "Olá mundo",
		"status":   "ok",
	})
}
func Veiculos(c *gin.Context) {

	var veiculos []models.Veiculo
	database.DB.Find(&veiculos)

	c.JSON(200, veiculos)
}
func GetVeiculo(c *gin.Context) {
	id := c.Params.ByName("id_veiculo")
	var veiculo models.Veiculo
	database.DB.First(&veiculo, id)

	if veiculo.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Erro": "Veiculo não encontrado",
		})
	} else {
		c.JSON(http.StatusOK, veiculo)
	}
}

func NovoVeiculo(c *gin.Context) {

	var veiculo models.Veiculo
	//validação json
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	} else { //validação dos campos do modelo
		if err := models.Validacao(&veiculo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
		} else {
			database.DB.Create(&veiculo)
			c.JSON(http.StatusOK, veiculo)
		}

	}

}

func DeletarVeiculo(c *gin.Context) {
	id := c.Params.ByName("id_veiculo")
	var veiculo models.Veiculo
	database.DB.Delete(&veiculo, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Veiculo deletado com sucesso",
	})
}

func AtualizarVeiculo(c *gin.Context) {
	id := c.Params.ByName("id_veiculo")
	var veiculo models.Veiculo
	database.DB.First(&veiculo, id)
	if err := c.ShouldBindJSON(&veiculo); err != nil { //validação json
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else { //validação dos campos do modelo
		if err := models.Validacao(&veiculo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			database.DB.Model(&veiculo).UpdateColumns(veiculo)
			c.JSON(http.StatusOK, gin.H{
				"message": "Veiculo atualizado com sucesso",
				"veiculo": veiculo,
			})
		}
	}

}
