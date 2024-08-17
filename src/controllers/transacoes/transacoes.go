package transacoes

import (
	"gorinha/src/commands"
	"gorinha/src/controllers/commom"
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
		commom.HandleDbHttpErrors(c, "Failed to update balance", err)
		return
	}

	if err := commands.UpdateBalance(id, payload.Value, payload.Type, payload.Description); err != nil {
		commom.HandleDbHttpErrors(c, "Failed to update balance", err)
		return
	}

	balances, _ := queries.GetBalance(id)

	c.JSON(http.StatusOK, balances)
}

// curl -X POST -H "Content-Type: application/json" -d '{ "valor": 100, "tipo": "c", "descricao": "descricao" }' localhost:9999/clientes/1/transacoes
