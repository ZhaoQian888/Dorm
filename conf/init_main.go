package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Start 初始化所有组间 并返回gin.Engine实例
func Start() *gin.Engine {
	godotenv.Load()
	db()
	redis()
	r := routerInit()
	return r
}
