package v1_controller

import (
	"encoding/json"
	"strconv"
	"vuejs-blog-server/controllers"
	"vuejs-blog-server/models"
	"vuejs-blog-server/request"
)

// 产品列表页
type UsersController struct {
	controllers.BasesController
}

func (this *UsersController) Prepare() {
	this.Validate(request.User{})
}

// 用户列表
func (this *UsersController) Index() {
	l, err := models.GetAllUsers(this.FileterParams())
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJSON()
}

// 创建用户
func (this *UsersController) Create() {
	this.Ctx.WriteString("Create")
}

// 保存用户
func (this *UsersController) Store() {
	var m models.Users

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &m)
	if err != nil {
		this.Data["json"] = this.ErrorResopnse(500, nil, err)
	} else {
		if _, err := models.AddUsers(&m); err != nil {
			this.Data["json"] = this.ErrorResopnse(500, nil, err)
		} else {
			this.Ctx.Output.SetStatus(201)
			this.Data["json"] = this.CreatedResopnse("用户创建成功")
		}
	}
	this.ServeJSON()
}

// 用户详情
func (this *UsersController) Show() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	user, err := models.GetUsersById(id)
	if err != nil {
		this.Data["json"] = this.ErrorResopnse(500, nil, err)
	} else {
		this.Data["json"] = this.SuccessResopnse(user, 200, nil)
	}
	this.ServeJSON()
}

// 编辑用户资料
func (this *UsersController) Edit() {
	this.Show()
}

// 更新用户资料
func (this *UsersController) Update() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	m := models.Users{
		Id: id,
	}
	if err := models.UpdateUsersById(&m); err != nil {
		this.Data["json"] = this.ErrorResopnse(500, nil, err)
	} else {
		this.Data["json"] = this.SuccessResopnse(200, nil, "资料更新成功")
	}
}

// 删除用户
func (this *UsersController) Destroy() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUsers(id); err != nil {
		this.Data["json"] = this.ErrorResopnse(500, nil, err)
	} else {
		this.Data["json"] = this.SuccessResopnse(200, nil, "删除成功")
	}
}
