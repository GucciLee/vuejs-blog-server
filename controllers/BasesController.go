package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"strings"
	"vuejs-blog-server/request"
	"vuejs-blog-server/utils/learnku_json"
)

const (
	HTTP_200 int = 200 // OK
	HTTP_201 int = 201 // Created
	HTTP_202 int = 202 // Accepted 认可的
	HTTP_203 int = 203 // Non-Authoritative Information 非授权信息
	HTTP_204 int = 204 // No Content 无内容
	HTTP_301 int = 301 // 永久重定向
	HTTP_302 int = 302 // 临时重定向
	HTTP_400 int = 400 // Bad Request 无效请求
	HTTP_401 int = 401 // Unauthorized 未授权请求
	HTTP_403 int = 403 // Forbidden 请求被禁止
	HTTP_404 int = 404 // Not Found 请求对象不存在
	HTTP_405 int = 405 // Method Not Allowed 方法不允许
	HTTP_500 int = 500 // Internal Server Error 内部服务器错误
	HTTP_501 int = 501 // Not Implemented 未实现
)

// @title 基础路由
type BasesController struct {
	beego.Controller
}

// @title 表单验证
// @params validFn 通过传入接口来实现调用
func (this BasesController) Validate(validFn request.RequestBaseI) {
	method := this.Ctx.Request.Method
	json2Map := map[string]interface{}{}
	b, err := learnku_json.Json2Map(this.Ctx.Input.RequestBody, &json2Map)
	if b {
		validErrorsData, validResult := validFn.Valid(method, json2Map)
		if !validResult {
			this.Data["json"] = this.ErrorResopnse(HTTP_200, nil, validErrorsData)
			this.ServeJSON()
		}
	} else {
		beego.Error(err)
	}
}

// @title 过滤请求参数
func (this BasesController) FileterParams() (map[string]string, []string, []string, []string, int64, int64) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	// where k = v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: 无效的查询键/值对")
				this.ServeJSON()
				return query, fields, sortby, order, offset, limit
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	return query, fields, sortby, order, offset, limit
}

// @title 201 创建操作
func (this BasesController) CreatedResopnse(message interface{}) interface{} {
	if message != nil {
		message = "创建成功"
	}
	return this.SuccessResopnse(nil, HTTP_201, message)
}

// @title 204 删除操作【没有内容的响应，一般用于 删除操作】
func (this BasesController) NoContentResopnse() interface{} {
	return this.SuccessResopnse(nil, HTTP_204, nil)
}

// @title 400 无效请求
func (this BasesController) ErrorBadResopnse() interface{} {
	message := map[string]string{
		"en-US": "Bad Request",
		"zh-CN": "无效请求",
	}
	return this.ErrorResopnse(HTTP_400, message, nil)
}

// @title 401 无权限访问
func (this BasesController) ErrorUnauthorizedResopnse() interface{} {
	message := map[string]string{
		"en-US": "Unauthorized",
		"zh-CN": "您无权限访问",
	}
	return this.ErrorResopnse(HTTP_401, message, nil)
}

// @title 403 禁止访问
func (this BasesController) ErrorForbiddenResopnse() interface{} {
	message := map[string]string{
		"en-US": "Forbidden",
		"zh-CN": "请求被禁止",
	}
	return this.ErrorResopnse(HTTP_403, message, nil)
}

// @title 404 页面
func (this BasesController) ErrorNoFoundResopnse() interface{} {
	message := map[string]string{
		"en-US": "Not Found",
		"zh-CN": "您所访问的页面不存在",
	}
	return this.ErrorResopnse(HTTP_404, message, nil)
}

// @title 405 方法不允许
func (this BasesController) ErrorMethodNotAllowedResopnse() interface{} {
	message := map[string]string{
		"en-US": "Method Not Allowed",
		"zh-CN": "您所请求的方法不被允许",
	}
	return this.ErrorResopnse(HTTP_405, message, nil)
}

// @title 500 内部服务器错误
func (this BasesController) ErrorInternalResopnse() interface{} {
	message := map[string]string{
		"en-US": "Internal Server Error",
		"zh-CN": "内部服务器错误，请联系管理员",
	}
	return this.ErrorResopnse(HTTP_500, message, nil)
}

// @title 成功返回数据格式
// @params data 成功数据
// @params code 默认200
// @params message 成功消息 可为空
func (this BasesController) SuccessResopnse(data, code, message interface{}) interface{} {
	if code == nil {
		code = HTTP_200
	}
	rtnJson := map[string]interface{}{}
	rtnJson["status"] = "success"
	rtnJson["code"] = code
	rtnJson["data"] = data
	rtnJson["message"] = message
	return rtnJson
}

// @title 失败返回数据格式
// @params code 默认404
// @params message 错误消息 与参数 error 至少保持一个
// @params errors 错误消息 与参数 message 至少保持一个
func (this BasesController) ErrorResopnse(code, message, errors interface{}) interface{} {
	if code == nil {
		code = HTTP_404
	}
	rtnJson := map[string]interface{}{}
	rtnJson["status"] = "error"
	rtnJson["code"] = code
	if errors != nil {
		rtnJson["error"] = errors
	} else {
		rtnJson["error"] = map[string]interface{}{
			"message": message,
		}
	}
	return rtnJson
}
