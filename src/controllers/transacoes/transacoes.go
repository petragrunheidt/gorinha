package transacoes

import (
	"gorinha/src/commands"
	"gorinha/src/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionPayload struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func HandleTransaction(c *gin.Context) {
	id := c.Param("id")

	var payload TransactionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := commands.UpdateBalance(id, payload.Value, payload.Type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balances, _ := queries.GetBalance(id)

	c.JSON(http.StatusOK, gin.H{
		"data": balances,
	})
}

// curl -X POST -H "Content-Type: application/json" -d '{ "valor": 1000, "tipo": "c", "descricao": "descricao" }' localhost:9999/clientes/1/transacoes
