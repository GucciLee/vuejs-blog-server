package request

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	RequestBaseS
}

func (this User) Valid(data map[string]interface{}) (map[string]string, bool) {
	valid := validation.Validation{}
	valid.Required(data["name"], "name").Message("用户名不能为空")
	valid.Required(data["email"], "email").Message("邮箱不能为空")

	return this.Rtn(valid)
}
