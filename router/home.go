package router

import (
	"gin-cli/controller"
	"github.com/gin-gonic/gin"
)

// 首页的Router

func RegisterHomeRouter(engine *gin.Engine) *gin.Engine {

	homerouter := engine.Group("/")
	{
		homerouter.GET("/", controller.MainController)

	}
	return engine
}
