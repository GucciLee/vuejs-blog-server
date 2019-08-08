package example

// Restful 接口
type RestfulController interface {
	// @Title 列表-页面
	// @Description 列表-页面
	// @Success 200 {object} models.Object
	// @router / [get]
	Index()
	// @Title 创建-页面
	// @Description 创建-页面
	// @Success 200 {object} models.Object
	// @router /create [get]
	Create()
	// @Title Store 保存
	// @Description Store 保存
	// @Param	body		body 	models.Object	true		"The object content"
	// @Success 200 {string} models.Object.Id
	// @Failure 403 body is empty
	// @router / [post]
	Store()
	// @Title Show 详情-页面
	// @Description 详情-页面
	// @Param	id		path 	string	true		"the id you want to get"
	// @Success 200 {object} models.Object
	// @Failure 403 :id is empty
	// @router /:id [get]
	Show()
	// @Title Edit 编辑-页面
	// @Description 编辑-页面
	// @Param	id		path 	string	true		"The id you want to update"
	// @Param	body		body 	models.Object	true		"The body"
	// @Success 200 {object} models.Object
	// @Failure 403 :id is empty
	// @router /:id/edit [get]
	Edit()
	// @Title Update 保存（更新）
	// @Description 保存（更新）
	// @Param	id		path 	string	true		"The id you want to update"
	// @Param	body		body 	models.Object	true		"The body"
	// @Success 200 {object} models.Object
	// @Failure 403 :id is empty
	// @router /:id [put,patch]
	Update()
	// @Title Delete 删除
	// @Description 删除
	// @Param	id		path 	string	true		"The id you want to delete"
	// @Success 200 {string} delete success!
	// @Failure 403 id is empty
	// @router /:id [delete]
	Destroy()
}