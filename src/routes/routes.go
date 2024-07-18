package routes

import (
	"gorinha/src/controllers/transacoes"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/clientes/:id/transacoes", transacoes.HandleTransaction)
	return r
}
