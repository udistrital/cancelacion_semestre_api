// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/juusechec/oas_be_cancelacion_semestre/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_cancelacion_semestre",
			beego.NSInclude(
				&controllers.TipoCancelacionSemestreController{},
			),
		),

		beego.NSNamespace("/estado_cancelacion_semestre",
			beego.NSInclude(
				&controllers.EstadoCancelacionSemestreController{},
			),
		),

		beego.NSNamespace("/cancelacion_semestre",
			beego.NSInclude(
				&controllers.CancelacionSemestreController{},
			),
		),

		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
