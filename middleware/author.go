package middleware

import (
	"Dorm/cache"
	"Dorm/serializer"
	"Dorm/service"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentOpenID 获取用户openid
func CurrentOpenID() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("loged")
		fmt.Println("中间件拿到的的loged", uid)
		a, _ := uid.(string)
		id, err := cache.GetOpenID(a)
		if err == nil && id != "" {
			c.Set("opid", id)
		}
		fmt.Print("中间件拿到的opid", id)
		c.Next()

	}
}

// NeedAuthor 需要登录
func NeedAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		opid, ok := c.Get("opid")
		if !ok {
			code := c.Query("code")
			fmt.Print("在中间件拿到的的code：", code)
			a, err := service.Login(code)
			if err != nil {
				c.JSON(200, serializer.BaseResponse{
					Status: 0055,
					Msg:    "获取openid失败",
					Data:   a,
				})
				c.Abort()
			}
			c.Set("wxid", a)
			c.Set("opid", a.OpenID)

			s := sessions.Default(c)
			s.Clear()
			s.Set("loged", a.OpenID)
			s.Save()

			r := cache.REDISPOOL.Get()
			opid, _ = c.Get("opid")
			opidinit, _ := opid.(string)
			r.Do("hset", opidinit, "openid", opidinit)
			r.Do("hset", opidinit, "sessionkey", a.SessionKey)
		}
		c.Next()
	}
}

// NeedStu 把所需要的学号查询并放到请求里
func NeedStu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("opid")
		i := id.(string)
		stunum, gen, err := service.GetStuNumGen(i)
		if err != nil {
			c.JSON(200, serializer.BaseResponse{
				Status: 332,
				Error:  err.Error(),
			})
		}
		c.Set("stu_number", stunum)
		c.Set("gender", gen)
	}
}
