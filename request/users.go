package request

import "fmt"

type Users struct {
	Model interface{}
}

func (this Users) Params() map[string]string {
	fmt.Println(this.Model)
	return map[string]string{
		"name": "a",
		"email": "a@qq.com",
	}
}

func (this Users) Rules() map[string]string {
	return map[string]string{
		"name": "required|min:6",
	}
}

func (this Users) Message() map[string]string {
	return map[string]string{
		"name.required": "用户名不能为空",
		"name.min": "用户名不少于6个字符",
	}
}

func (this Users) Go() {
	this.Params()
	// isok := Validation{this.model}
	// isok.Begin(this.Params(),this.Rules(),this.Message())
}