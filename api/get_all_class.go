package api

import (
	"Dorm/serializer"

	"github.com/gin-gonic/gin"
)

// Classes 所有学苑列表。无需查询数据库
var Classes = []string{
	"求知一苑",
	"求知二苑",
	"求知三苑",
}

// AllClasses 所有学苑列表
func AllClasses(c *gin.Context) {
	c.JSON(200, serializer.BuildListResponse(Classes, len(Classes)))
}
