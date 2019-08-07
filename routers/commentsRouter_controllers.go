package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put","patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Destroy",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/:id/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:ExampleController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"] = append(beego.GlobalControllerRouter["vuejs-blog-server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
