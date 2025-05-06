package user

import (
	"fmt"
)

type LoginRequest struct {
	Telephone    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (r *LoginRequest) Examine() error {
	telephone := r.Telephone
	password := r.Password
	if telephone == "" || password == "" {
		return fmt.Errorf("手机号和密码不能为空")
	}
	return nil
}

type LoginResponse struct {
	UserInfo
}