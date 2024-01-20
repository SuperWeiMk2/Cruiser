package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 根路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	// 用户路由
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/info", func(ctx *gin.Context) {
			ctx.String(200, "User Info")
		})
	}
	// 数据库路由
	dbRoutes := r.Group("/db")
	{
		dbRoutes.GET("/info", func(ctx *gin.Context) {
			ctx.String(200, "db Info")
		})
	}
	// 监控路由
	monitorRoutes := r.Group("/monitor")
	{
		monitorRoutes.GET("/info", func(ctx *gin.Context) {
			ctx.String(200, "monitor Info")
		})
		monitorRoutes.GET("/index", func(ctx *gin.Context) {

		})
	}
}
