package utils

import "fmt"

var (
	ErrEmailExist        = fmt.Errorf("email already exists")                          // 邮箱已被注册
	ErrCaptcha           = fmt.Errorf("invalid captcha")                               // 验证码错误
	ErrIncorrectPassword = fmt.Errorf("incorrect password")                            // 密码错误
	ErrUserNotFound      = fmt.Errorf("user not found")                                // 用户未找到
	ErrShortLinkExist    = fmt.Errorf("short link already exists")                     // 短链接已存在
	ErrNoShortLink       = fmt.Errorf("short link does not exist")                     // 短链接不存在
	ErrPrivilege         = fmt.Errorf("insufficient privileges")                       // 权限不足
	ErrShortLinkActive   = fmt.Errorf("short link is currently active")                // 短链接正在使用
	ErrShortLinkTime     = fmt.Errorf("short link is not within the valid time range") // 短链接不在时间范围内
)

func NewCustomError(message string) error {
	return fmt.Errorf("custom error: %s", message)
}
