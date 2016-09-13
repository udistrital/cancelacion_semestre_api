package main

import (
	_ "github.com/JorgeUlises/oas_be_cancelacion_semestre/routers"

	"github.com/astaxie/beego"
<<<<<<< HEAD
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://cancelacion_semestre:docker@localhost:5432/udistrital?sslmode=disable")
}

=======
)

>>>>>>> 9e7d11bd4d786705bcdcfd3f9612c17388e95716
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
<<<<<<< HEAD

=======
>>>>>>> 9e7d11bd4d786705bcdcfd3f9612c17388e95716
