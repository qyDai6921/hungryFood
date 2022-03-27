package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"workspace/ginweb/conf"
	"workspace/ginweb/handler"
)

func RouterRun(port int, args ...string) error {

	gin.SetMode(gin.ReleaseMode)

	// router
	router := gin.Default()
	router.Use(Middelware())

	// Front-end router
	router.GET("/api/search_menus", handler.SearchMenus) // search_menus
	router.POST("/api/order", handler.Order)             // order

	return router.Run(fmt.Sprintf(":%d", port))
}

//路由中间件
func Middelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Length,Content-Type,X-Token,Access-Token,Token,Sign,Tm,UUID,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requeste")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		//nginx 重复了, 上生产注释掉
		if conf.RunMode == "debug" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
