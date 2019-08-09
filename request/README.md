## 使用案例
在控制器文件中使用

### request 表单验证器
```
type User struct {
	RequestBaseS
}

func (this User) Valid(data map[string]interface{}) (map[string]string, bool) {
	valid := validation.Validation{}
	valid.Required(data["name"], "name").Message("名称不能为空")
	return this.Rtn(valid)
}
```

### controller 控制器
```
func (c *UsersController) Prepare() {
	json2Map := learnku_json.Json2Map(c.Ctx.Input.RequestBody)
	valid, b := request.User{}.Valid(json2Map)
	// 验证不通过，响应错误信息
	if !b {
		c.Data["json"] = valid
		c.ServeJSON()
	}
}
```