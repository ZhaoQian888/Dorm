package conf

import (
	"Dorm/api"

	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", api.Ping)
	router.GET("/dorm/remain", api.RemainDorm)
	router.POST("/distr", api.DistriDorm)
	return router
}
