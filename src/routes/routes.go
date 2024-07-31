package routes

import (
	"gorinha/src/controllers/transacoes"
	"gorinha/src/controllers/extrato"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/clientes/:id/extrato", extrato.HandleExtract)
	r.POST("/clientes/:id/transacoes", transacoes.HandleTransaction)
	return r
}
