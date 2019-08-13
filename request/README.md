## 使用案例
[Beego 表单验证](https://beego.me/docs/mvc/controller/validation.md)

### request 表单验证器
```
type User struct {
	RequestBaseS
}

// 定义表单验证规则
func (this User) Valid(data map[string]interface{}) (map[string]string, bool) {
	valid := validation.Validation{}
    
    // todu...
	valid.Required(data["name"], "name").Message("名称不能为空")
    // 自定义验证规则
	if true {
		_ = valid.SetError("ruler", "我是自定义验证规则")
	}

	return this.Rtn(valid)
}
```

### controller 控制器
```
func (c *UsersController) Prepare() {
	// 表单验证
	c.Validate(request.User{})
}
```