package helpers

import (
	"gorinha/src/db"

	"github.com/gin-gonic/gin"
)

func InitTestDB() {
	gin.SetMode(gin.TestMode)
	db.Init()
	db.Drop()
	db.Migrate()
}
