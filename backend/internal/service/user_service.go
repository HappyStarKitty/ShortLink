// user service
package service

import (
	"backend/api/dto"
	"backend/internal/dao/model"
	"backend/utils"
	"errors"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserService interface {
	RegisterUser(req dto.RegisterUserRequest) (uint, error)
	LoginUser(req dto.LoginUserRequest) (string, error)
	LogoutUser(c *gin.Context) error
	GetUserInfo(userID uint) (*dto.UserInfoResponse, error)
	UpdateUserInfo(userID uint, req dto.UpdateUserInfoRequest) error
	UpdatePassword(userID uint, req dto.UpdatePasswordRequest) error
	GetCaptcha() (string, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// 用户注册
func (s *userService) RegisterUser(req dto.RegisterUserRequest) (uint, error) {
	var user model.User
	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return 0, utils.ErrEmailExist
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	newUser := model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(&newUser).Error; err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

// 用户登入
func (s *userService) LoginUser(req dto.LoginUserRequest) (string, error) {
	var user model.User

	// 查询用户
	err := s.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		if err == gorm.ErrRecordNotFound {
			return "", utils.ErrUserNotFound
		}
		return "", err
	}

	// 打印用户信息
	log.Printf("User found: %+v", user)

	// 检查密码
	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		log.Printf("Password check failed for user ID: %d", user.ID)
		return "", utils.ErrIncorrectPassword
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return "", err
	}

	return token, nil
}

// 用户登出
func (s *userService) LogoutUser(c *gin.Context) error {
	token := c.GetHeader("Authorization")
	if token == "" {
		return errors.New("no token provided")
	}

	return nil
}

// 获取用户信息
func (s *userService) GetUserInfo(userID uint) (*dto.UserInfoResponse, error) {
	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	userInfo := &dto.UserInfoResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userInfo, nil
}

// 更新用户信息
func (s *userService) UpdateUserInfo(userID uint, req dto.UpdateUserInfoRequest) error {
	if err := s.db.Model(&model.User{}).Where("id = ?", userID).Update("name", req.Name).Error; err != nil {
		return err
	}

	return nil
}

// 更新密码
func (s *userService) UpdatePassword(userID uint, req dto.UpdatePasswordRequest) error {
	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := utils.CheckPassword(req.OldPassword, user.Password); err != nil {
		return utils.ErrIncorrectPassword
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	if err := s.db.Model(&user).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}

// 生成验证码并返回 captcha_id
func (s *userService) GetCaptcha() (string, error) {
	captchaID := captcha.New()
	return captchaID, nil
}
