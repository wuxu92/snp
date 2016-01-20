package routers

import (
	"snp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/mongo", &controllers.MongoController{})
    beego.Router("/init", &controllers.InitController{})
    beego.Router("/api/:model/:action/:id", &controllers.ApiController{})
}
