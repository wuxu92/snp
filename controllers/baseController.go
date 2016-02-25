package controllers
import "github.com/astaxie/beego"

type BaseController struct {
  beego.Controller
}

func (this *BaseController) Json(data ResObj) {
  this.Data["json"] = data.Json()
  this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf8")
  this.ServeJson()
}