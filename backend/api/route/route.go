package route

import (
	"backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, linkController controller.LinkController, userController controller.UserController) {
	// link 路由注册
	r.POST("/api/link/create", linkController.CreateLink)         // 创建短链接
	r.GET("/:shortCode", linkController.GetLink)                  // 获取短链接信息
	r.PUT("/api/link/info/:shortCode", linkController.UpdateLink) // 更新短链接信息
	r.POST("/api/link/delete", linkController.DeleteLink)         // 删除短链接
	r.GET("/api/link/list", linkController.ListLinks)             // 获取短链接列表

	// user 路由注册
	r.POST("/api/user/register", userController.RegisterUser)  // 用户注册
	r.POST("/api/user/login", userController.LoginUser)        // 用户登录
	r.POST("/api/user/logout", userController.LogoutUser)      // 用户登出
	r.GET("/api/user/info", userController.GetUserInfo)        // 获取用户信息
	r.PUT("/api/user/info", userController.UpdateUserInfo)     // 修改用户信息
	r.PUT("/api/user/password", userController.UpdatePassword) // 修改密码

	// captcha 路由注册
	r.GET("/api/user/captcha", userController.GetCaptcha)                    // 获取登录验证码
	r.GET("/api/user/captcha/:captcha_id", userController.ServeCaptchaImage) // 获取登录验证码

}
