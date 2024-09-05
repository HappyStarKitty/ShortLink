// user controller
package controller

import (
	"backend/api/dto"
	"backend/internal/service"
	"backend/utils"
	"encoding/json"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
	GetUserInfo(c *gin.Context)
	UpdateUserInfo(c *gin.Context)
	UpdatePassword(c *gin.Context)
	GetCaptcha(c *gin.Context)
	ServeCaptchaImage(c *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

// 用户注册
func (ctrl *userController) RegisterUser(c *gin.Context) {
	var req dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("RegisterUser error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ctrl.service.RegisterUser(req)
	if err != nil {
		log.Printf("RegisterUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("RegisterUser success: userID=%d", userID)

	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

// 用户登入
func (ctrl *userController) LoginUser(c *gin.Context) {
	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("LoginUser bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.LoginUser(req)
	if err != nil {
		if err == utils.ErrUserNotFound || err == utils.ErrIncorrectPassword {
			log.Printf("LoginUser invalid credentials: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			log.Printf("LoginUser error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Printf("LoginUser success: token=%s", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// 用户登出
func (ctrl *userController) LogoutUser(c *gin.Context) {
	if err := ctrl.service.LogoutUser(c); err != nil {
		log.Printf("LogoutUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("LogoutUser success")
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// 获取用户信息
func (ctrl *userController) GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		log.Printf("GetUserInfo unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userInfo, err := ctrl.service.GetUserInfo(userID.(uint))
	if err != nil {
		log.Printf("GetUserInfo error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetUserInfo success: user=%v", userInfo)
	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}

// 更新用户信息
func (ctrl *userController) UpdateUserInfo(c *gin.Context) {
	var req dto.UpdateUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("UpdateUserInfo bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		log.Printf("UpdateUserInfo unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := ctrl.service.UpdateUserInfo(userID.(uint), req); err != nil {
		log.Printf("UpdateUserInfo error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateUserInfo success")
	c.JSON(http.StatusOK, gin.H{"message": "User info updated successfully"})
}

// 修改密码
func (ctrl *userController) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("UpdatePassword bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		log.Printf("UpdatePassword unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := ctrl.service.UpdatePassword(userID.(uint), req); err != nil {
		log.Printf("UpdatePassword error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdatePassword success")
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// 获取验证码
func (ctrl *userController) GetCaptcha(c *gin.Context) {
	captchaID, err := ctrl.service.GetCaptcha()
	if err != nil {
		log.Printf("GetCaptcha error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成验证码URL
	captchaURL := "http://" + c.Request.Host + "/api/user/captcha/" + captchaID
	// captcha response
	response := gin.H{"captcha_id": captchaID, "captcha_url": captchaURL}
	responseJSON, _ := json.Marshal(response)
	log.Printf("GetCaptcha success: %s", responseJSON)
	c.JSON(http.StatusOK, response)
}

func (ctrl *userController) ServeCaptchaImage(c *gin.Context) {
	captchaID := c.Param("captcha_id")

	c.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(c.Writer, captchaID, 240, 80); err != nil {
		log.Printf("ServeCaptchaImage error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serve captcha image"})
		return
	}

	log.Printf("ServeCaptchaImage success for captcha_id=%s", captchaID)
}
