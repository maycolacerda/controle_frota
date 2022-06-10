package routes

import (
	"controle_frota_gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.Index)

	http.ListenAndServe("localhost:8000", r)

}
