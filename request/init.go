package request

import "github.com/astaxie/beego/validation"

type OverflowValidation struct {
}

// 设置默认的表单验证提示信息
func (this OverflowValidation) SetDefaultMessage()  {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "不能为空2",
		"Min":          "Minimum is %d",
		"Max":          "Maximum is %d",
		"Range":        "Range is %d to %d",
		"MinSize":      "Minimum size is %d",
		"MaxSize":      "Maximum size is %d",
		"Length":       "Required length is %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	})
} 

// 表单验证
func init()  {
	self := OverflowValidation{}
	self.SetDefaultMessage()
}
