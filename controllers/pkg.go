package controllers

import (
	"fmt"
	"snp/models"
	//	"strings"

	"github.com/astaxie/beego"
	"time"
)

type PkgController struct {
	beego.Controller
}

func (this *PkgController) Get() {
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	fmt.Println("model: pkg", " action: ", action, "id: ", id)

	switch action {
	case "get":
		if id == "d" {
			id = "default"
		}
		fmt.Println("getting pkg:", id)
		this.Data["json"] = models.GetPkgFullInfo(id)
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf8")
		this.ServeJson()
	case "new":
	case "fork":
		fmt.Println("fork pkg:", id)
		if id == "" {
			this.Data["json"] = false
		} else {
			name := this.GetString("name", "pkg-"+time.Now().Format(time.Stamp))
			pkg, err := forkPkg(id, name)
			data := make(map[string]interface{})
			if err != nil {
				data["code"] = -1
				data["err"] = err
				this.Data["json"] = &data
			} else {
				data["code"] = 0
				data["err"] = pkg.Name
				this.Data["json"] = &data
			}
		}
		this.ServeJson()
	}

}

func forkPkg(name string, newName string) (models.Pkg, error){
	// copy package
	old, err := models.GetPkgByName(name)
	if err != nil {
		return models.Pkg{}, err
	}
	return old.Copy(newName)
}

func NewPkg() {

}
