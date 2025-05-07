package user

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

type SendVCodeRequest struct {
	Telephone string `json:"telephone" binding:"required,telephone"`
}

func (r *SendVCodeRequest) Examine() error {
	reg := regexp2.MustCompile(`/^(?:(?:\+|00)86)?1\d{10}$/`, 0)
	isMatch, _ := reg.MatchString(r.Telephone)
	if !isMatch {
		return fmt.Errorf("邮箱不合格")
	}

	return nil
}
