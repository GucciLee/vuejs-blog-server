package example

import "github.com/astaxie/beego"

// 产品列表页
type ExampleController struct {
	beego.Controller
}

// @Title 列表 - 展示
// @Description 列表 - 展示
// @Success 200 {object} models.Object
// @router / [get]
func (this *ExampleController) Index()  {
	this.Ctx.WriteString("Get - 列表 - 展示")
}

// @Title 创建 - 展示
// @Description 创建 - 展示
// @Success 200 {object} models.Object
// @router /create [get]
func (this *ExampleController) Create()  {
	this.Ctx.WriteString("Create - 创建 - 展示")
}

// @Title Store 保存
// @Description Store 保存
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ExampleController) Store()  {
	this.Ctx.WriteString("Store")
}

// @Title Show 详情 - 展示
// @Description 详情 - 展示
// @Param	id		path 	string	true		"the id you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id [get]
func (this *ExampleController) Show()  {
	this.Ctx.WriteString("Show")
}

// @Title Edit 编辑 - 展示
// @Description 编辑 - 展示
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id/edit [get]
func (this *ExampleController) Edit()  {
	this.Ctx.WriteString("Edit")
}

// @Title Update 更新
// @Description 更新
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id [put,patch]
func (this *ExampleController) Update()  {
	this.Ctx.WriteString("Update")
}

// @Title Delete 删除
// @Description 删除
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ExampleController) Destroy()  {
	this.Ctx.WriteString("Destroy")
}