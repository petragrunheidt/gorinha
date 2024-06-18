package routes

import (
	"gorinha/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", controllers.Ping)
    return r
}