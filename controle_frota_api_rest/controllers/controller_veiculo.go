package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// It returns a JSON response with a 200 status code
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
		"status":  "ok",
	})
}

// It gets all the vehicles from the database and returns them in JSON format
func Veiculos(c *gin.Context) {

	var veiculos []models.Veiculo
	database.DB.Find(&veiculos)

	c.JSON(200, veiculos)
}

// It gets the id from the URL, then it looks for the veiculo in the database, if it doesn't find it,
// it returns a 404 error, if it finds it, it returns the veiculo
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

// If the JSON is valid, then validate the model, if the model is valid, then create the record in the
// database
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

// It deletes a vehicle from the database
func DeletarVeiculo(c *gin.Context) {
	id := c.Params.ByName("id_veiculo")
	var veiculo models.Veiculo
	database.DB.Delete(&veiculo, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Veiculo deletado com sucesso",
	})
}

// It receives a JSON object, validates it, and then updates the database with the new values
func AtualizarVeiculo(c *gin.Context) {
	id := c.Params.ByName("id_veiculo")
	var veiculo models.Veiculo
	database.DB.First(&veiculo, id)
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	err := models.Validacao(&veiculo)
	if err != nil {
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

// It renders the index.html template, passing in the Message variable, which is set to "Boas vindas"
func RenderIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "Boas vindas",
	})
}
