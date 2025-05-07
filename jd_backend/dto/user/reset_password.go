package user

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

type ResetPasswordRequest struct {
	Telephone string `json:"telephone"`
	VCode     string `json:"vcode"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

func (r *ResetPasswordRequest) Examine() error {
	if r.Password != r.Password2 {
		return fmt.Errorf("两次密码输入不一致")
	}
	reg := regexp2.MustCompile(`^\S*(?=\S{6,})(?=\S*\d)(?=\S*[A-Z])(?=\S*[a-z])(?=\S*[!@#$%^&*? ])\S*$`, 0)
	isMatch, _ := reg.MatchString(r.Password)
	if !isMatch {
		return fmt.Errorf("密码必须包含数字、大写字母、小写字母和特殊字符，且长度至少为6位")
	}
	// 去除空格查看是否为空
	// 去除空格
	telephone := strings.TrimSpace(r.Telephone)
	vcode := strings.TrimSpace(r.VCode)
	if telephone == "" || vcode == "" {
		return fmt.Errorf("用户名、邮箱和验证码不能为空")
	}
	if len(telephone) < 5 || len(telephone) > 50 {
		return fmt.Errorf("手机号度必须在5到50个字符之间")
	}

	return nil
}
