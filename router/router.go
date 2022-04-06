package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"workspace/ginweb/conf"
	"workspace/ginweb/handler"
)

func RouterRun(port int, args ...string) error {

	//在上线的时候，一定要选择release模式
	gin.SetMode(gin.ReleaseMode)

	// router
	router := gin.Default()  // gin.Default() 函数会生成一个默认的 Engine 对象:Logger 和 Recovery
	router.Use(Middelware()) // (这行可以去掉) router.use() — router 对象创建后，就可以像应用一样添加中间件和HTTP方法（get、put、post等）了

	// Front-end router: API url
	router.GET("/api/search_menus", handler.SearchMenus) // search_menus
	router.POST("/api/order", handler.Order)             // order

	return router.Run(fmt.Sprintf(":%d", port)) // fmt.Sprintf 格式化字符串
	// 控制台输出：API server listening at: [::]:59669
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
