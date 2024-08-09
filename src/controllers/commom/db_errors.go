package commom

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleDbHttpErrors(c *gin.Context, m string, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": m})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("%s: %v", m, err)})
	}
}