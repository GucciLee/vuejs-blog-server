package example

import "github.com/astaxie/beego"

// 产品列表页
type ExampleController struct {
	beego.Controller
}

// @Title 列表 - 展示
// @router / [get]
func (this *ExampleController) Index() {
	this.Ctx.WriteString("Get - 列表 - 展示")
}

// @Title 创建 - 展示
// @router /create [get]
func (this *ExampleController) Create() {
	this.Ctx.WriteString("Create - 创建 - 展示")
}

// @Title Store 保存
// @router / [post]
func (this *ExampleController) Store() {
	this.Ctx.WriteString("Store")
}

// @Title Show 详情 - 展示
// @router /:id [get]
func (this *ExampleController) Show() {
	this.Ctx.WriteString("Show")
}

// @Title Edit 编辑 - 展示
// @router /:id/edit [get]
func (this *ExampleController) Edit() {
	this.Ctx.WriteString("Edit")
}

// @Title Update 更新
// @router /:id [put,patch]
func (this *ExampleController) Update() {
	this.Ctx.WriteString("Update")
}

// @Title Delete 删除
// @router /:id [delete]
func (this *ExampleController) Destroy() {
	this.Ctx.WriteString("Destroy")
}
