package api

import (
	"backend/api/route"
	"backend/internal/controller"
	"backend/internal/dao"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

func InitAPI() *gin.Engine {
	r := gin.Default()

	// 初始化数据库
	db, err := dao.InitDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// 初始化服务
	linkService := service.NewLinkService(dao.NewLinkDAO(db))
	userService := service.NewUserService(db)

	// 初始化控制器
	linkController := controller.NewLinkController(linkService)
	userController := controller.NewUserController(userService)

	// 控制器注册路由
	route.RegisterRoutes(r, linkController, userController)

	return r
}
