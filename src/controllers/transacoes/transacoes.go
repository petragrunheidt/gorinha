package transacoes

import (
	"encoding/json"
	"fmt"
	"gorinha/src/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	query := getBalanceQuery("1")
	result, _ := db.ExecuteQuery(query)
	jsonResponse, _ := json.Marshal(result)
	
	c.JSON(http.StatusOK, gin.H{
		"message": jsonResponse,
	})
}

func getBalanceQuery(id string) string {
	return fmt.Sprintf("SELECT a.limit_amount, b.amount FROM accounts AS a JOIN balances AS b ON a.id = b.account_id WHERE a.id = '%s'", id)
}
