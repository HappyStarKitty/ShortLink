package dto

import (
// "github.com/guregu/null"
)

type LoginReq struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	CaptchaID       string `json:"captcha_id" binding:"required"`
	CaptchaSolution string `json:"captcha_solution" binding:"required"`
}

type LoginResp struct {
	UserID string `json:"user_id"`
}
