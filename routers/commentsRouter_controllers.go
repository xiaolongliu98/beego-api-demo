package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"] = append(beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"],
		beego.ControllerComments{
			Method:           "GetUserById",
			Router:           "/id/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"] = append(beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"],
		beego.ControllerComments{
			Method:           "GetIntroById",
			Router:           "/intro/id/:iid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"] = append(beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"],
		beego.ControllerComments{
			Method:           "GetIntroByUid",
			Router:           "/intro/uid/:uid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"] = append(beego.GlobalControllerRouter["bee-api-example1/controllers/user:UserController"],
		beego.ControllerComments{
			Method:           "GetUserByUsername",
			Router:           "/username/:username",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
