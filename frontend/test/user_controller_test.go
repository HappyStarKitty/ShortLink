package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/api/dto"
	"backend/internal/controller"
	"backend/internal/dao"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupUserRouter(userCtrl *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.POST("/api/user/register", userCtrl.RegisterUser)
	r.POST("/api/user/login", userCtrl.LoginUser)
	return r
}

func TestRegisterUser(t *testing.T) {
	db := setupTestDB(t)
	userDAO := dao.NewUserDAO(db)
	userService := service.NewUserService(userDAO)
	userCtrl := controller.NewUserController(userService)
	router := setupUserRouter(userCtrl)

	registerRequest := dto.RegisterUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(registerRequest)

	req, _ := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user_id")
}

func TestLoginUser(t *testing.T) {
	db := setupTestDB(t)
	userDAO := dao.NewUserDAO(db)
	userService := service.NewUserService(userDAO)
	userCtrl := controller.NewUserController(userService)
	router := setupUserRouter(userCtrl)

	user := dao.NewUserDAO(db)
	user.RegisterUser(dto.RegisterUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	})

	loginRequest := dto.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(loginRequest)

	req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
