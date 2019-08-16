package v1_controller

import "vuejs-blog-server/controllers"

// DemosController operations for Demos
type DemosController struct {
	controllers.BasesController
}

// URLMapping ...
func (c *DemosController) URLMapping() {
	// c.Mapping("Index", c.Index)
}

func (c *DemosController) Index() {
	c.Ctx.WriteString("Index")
}

func (c *DemosController) Create() {
	c.Ctx.WriteString("Create")
}

func (c *DemosController) Store() {
	c.Ctx.WriteString("Store")
}

func (c *DemosController) Show() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("Show" + id)
}

func (c *DemosController) Edit() {
	c.Ctx.WriteString("Edit")
}

func (c *DemosController) Update() {
	c.Ctx.WriteString("Update")
}

func (c *DemosController) Destroy() {
	c.Ctx.WriteString("Destroy")
}