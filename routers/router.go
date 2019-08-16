// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"strings"
	v1 "vuejs-blog-server/controllers/v1_controller"
)

func init() {
	Restfull("/v1/test", &v1.DemosController{}, nil)
}

func Restfull(rootpath string, c beego.ControllerInterface, methods []string)  {
	if len(methods) <= 0 {
		methods = []string{"Index", "Create", "Store", "Show", "Edit", "Update", "Destroy"}
	}
	for _, v := range methods{
		switch strings.ToLower(v) {
		case "index":
			beego.Router(rootpath + "/", c, "get:Index")
		case "create":
			beego.Router(rootpath + "/create", c, "get:Create")
		case "store":
			beego.Router(rootpath + "/:id", c, "post:Store")
		case "show":
			beego.Router(rootpath + "/:id", c, "get:Show")
		case "edit":
			beego.Router(rootpath + "/:id/edit", c, "get:Edit")
		case "update":
			beego.Router(rootpath + "/:id", c, "put,patch:Update")
		case "destroy":
			beego.Router(rootpath + "/:id", c, "delete:Destroy")
		default:
			beego.Router(rootpath + "/", c, "get:Index")
		}
	}
}
