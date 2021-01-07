package base

import (
	"gin-cli/router"
	"gin-cli/util"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterEngine() *gin.Engine {
	engine := gin.Default()
	engine.Static("/static", "static")
	engine.LoadHTMLGlob("view/**/*")
	engine = router.RegisterRouter(engine)
	InitDataBase()	// 初始化数据库，如果尚未添加可以先注释掉
	return engine
}

func RunEngine(engine *gin.Engine) {
	path := util.GetConfig("server", "addr")
	if err := engine.Run(path); err != nil {
		log.Fatal(err)
	}
}