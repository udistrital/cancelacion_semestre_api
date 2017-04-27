package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:CancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:EstadoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:TipoCancelacionSemestreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oas_be_cancelacion_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/getToken`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
