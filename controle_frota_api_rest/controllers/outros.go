package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {

	c.JSON(http.StatusNotFound, gin.H{

		"Erro": "Página não encontrada",
	})

}
