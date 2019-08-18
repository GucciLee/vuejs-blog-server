package request

import (
	"github.com/astaxie/beego/validation"
	"vuejs-blog-server/models"
)

type User struct {
	RequestBaseS
}

func (this User) Valid(method string, data map[string]interface{}) (map[string]string, bool) {
	valid := validation.Validation{}
	if method == "POST" {
		if models.ExistsUserName(data["name"]) {
			valid.AddError("name", "用户名已经存在")
		}
		valid.MaxSize(data["name"], 25, "name").Message("用户名必须介于 2 - 25 个字符之间。")
		valid.MinSize(data["name"], 2, "name").Message("用户名必须介于 2 - 25 个字符之间。")
		valid.Required(data["name"], "name").Message("用户名不能为空")

		if models.ExistsUserName(data["email"]) {
			valid.AddError("email", "邮箱已经存在")
		}
		valid.Email(data["email"], "email").Message("请填写正确的邮箱")
		valid.Required(data["email"], "email").Message("邮箱不能为空")

		valid.MinSize(data["password"],6, "password").Message("密码长度必须介于 2 - 25 个字符之间。")
		valid.MaxSize(data["password"],25, "password").Message("密码长度必须介于 2 - 25 个字符之间。")
	}
	return this.Rtn(valid)
}
