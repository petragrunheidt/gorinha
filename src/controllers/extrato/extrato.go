package extrato

import (
	"fmt"
	"gorinha/src/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountExtract struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func HandleExtract(c *gin.Context) {
	id := c.Param("id")

	extract, err := queries.GetExtract(id)
	if err != nil {
		fmt.Printf("Failed to get extract for ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve extract"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": extract,
	})
}
