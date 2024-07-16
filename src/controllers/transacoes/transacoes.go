package transacoes

import (
	"gorinha/src/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id := c.Param("id")
	balances, _ := queries.GetBalance(id)
	
	c.JSON(http.StatusOK, gin.H{
		"data": balances,
	})
}
