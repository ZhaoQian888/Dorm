package conf

import (
	"Dorm/api"
	"Dorm/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", api.Ping)
	router.GET("/wx/classes", api.AllClasses)

	router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.Cors())

	router.GET("/wx/bed/toremain", api.ToRemainBed)

	router.Use(middleware.CurrentOpenID())
	router.Use(middleware.NeedAuthor())

	router.GET("/wx/myid", api.ReturnID)
	router.GET("/wx/myinfo", api.GetInfo)

	router.POST("/wx/bind/info", api.BindInfo)

	router.GET("/wx/admin/push", api.AdminPush)
	router.GET("wx/wthbind/", api.WthBind)
	router.Use(middleware.NeedStu())
	router.POST("/wx/distr", api.DistriDorm)
	router.GET("wx/wtcho", api.WtChoose)
	return router
}
