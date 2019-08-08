package request

// 表单验证
type Validation struct {
	Model interface{}
}

// 开始校验
func (this Validation) Begin(params, rules, message map[string]string) {

}

/*func (this Validation) Rules() map[string]string {
	rules := map[string]string{
		"name": "required",
	}
	return rules
}

func (this Validation) Message() map[string]string {
	message := map[string]string{
		"name.required": "xxx不能为空",
	}
	return message
}*/