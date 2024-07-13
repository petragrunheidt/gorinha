package transacoes

import (
	"gorinha/src/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	balances, _ := queries.GetBalance(1)
	
	c.JSON(http.StatusOK, gin.H{
		"data": balances,
	})
}
