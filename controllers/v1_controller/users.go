package v1_controller

import (
	"errors"
	"strings"
	"vuejs-blog-server/controllers"
	"vuejs-blog-server/models"
)

// 产品列表页
type UsersController struct {
	controllers.BasesController
}

func (c *UsersController) Index() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (this *UsersController) Create() {
	this.Ctx.WriteString("Create - 创建 - 展示")
}

func (this *UsersController) Store() {
	this.Ctx.WriteString("Store")
}


func (this *UsersController) Show() {
	this.Ctx.WriteString("Show")
}


func (this *UsersController) Edit() {
	this.Ctx.WriteString("Edit")
}


func (this *UsersController) Update() {
	this.Ctx.WriteString("Update")
}

func (this *UsersController) Destroy() {
	this.Ctx.WriteString("Destroy")
}
