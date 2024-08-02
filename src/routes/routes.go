package routes

import (
	"gorinha/src/controllers/extrato"
	"gorinha/src/controllers/transacoes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/clientes/:id/extrato", extrato.HandleExtract)
	r.POST("/clientes/:id/transacoes", transacoes.HandleTransaction)

	r.NoRoute(handleNotFound)

	return r
}

func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
}