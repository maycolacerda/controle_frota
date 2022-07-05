package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// If the requested URL is not found, return a JSON response with the error message "Page not found"
func NotFound(c *gin.Context) {

	c.JSON(http.StatusNotFound, gin.H{

		"Erro": "Página não encontrada",
	})

}
