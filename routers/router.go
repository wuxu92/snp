package routers

import (
	"snp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/mongo", &controllers.MongoController{})
	beego.Router("/init", &controllers.InitController{})
	// action : get/update/delete
	// id: in fact it is name
	beego.Router("/api/pkg/:action/?:id", &controllers.PkgController{})
	beego.Router("/api/grp/:action/?:id", &controllers.GrpController{})
	beego.Router("/api/site/:action/?:id", &controllers.SiteController{})
	beego.Router("/api/new/:model", &controllers.NewModelController{})
	// beego.Router("/api/:model/:action/:id", &controllers.ApiController{})
}
