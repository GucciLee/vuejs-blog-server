package request

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	RequestBaseS
}

func (this User) Valid(data map[string]interface{}) (map[string]string, bool) {
	valid := validation.Validation{}
	valid.Required(data["name"], "name").Message("名称不能为空")
	return this.Rtn(valid)
}
