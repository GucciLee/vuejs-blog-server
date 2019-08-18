package v1_controller

import (
	"vuejs-blog-server/controllers"
)

// DemosController operations for Demos
type DemosController struct {
	controllers.BasesController
}

func (c *DemosController) Prepare() {
	// 表单验证
	// c.Validate(request.User{})
}

// 用户列表
func (c *DemosController) Index() {

}

// 注册
func (c *DemosController) Create() {
	c.Ctx.WriteString("Create")
}

// 注册
func (c *DemosController) Store() {

}

// 个人中心
func (c *DemosController) Show() {

}

// 编辑个人资料
func (c *DemosController) Edit() {
	c.Ctx.WriteString("Edit")
}

// 编辑个人资料
func (c *DemosController) Update() {
	c.Ctx.WriteString("Update")
}

// 删除用户
func (c *DemosController) Destroy() {
	c.Ctx.WriteString("Destroy")
}