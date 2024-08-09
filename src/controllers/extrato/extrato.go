package extrato

import (
	"gorinha/src/controllers/commom"
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
		commom.HandleDbHttpErrors(c, "Failed to retrieve extract", err)
		return
	}

	c.JSON(http.StatusOK, extract)
}
