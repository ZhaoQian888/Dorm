package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Init 初始化所有组间 并返回gin.Engine实例
func Init() *gin.Engine {
	godotenv.Load()
	db()
	redis()
	r := routerInit()
	return r
}
