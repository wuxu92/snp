package controllers

import (
	// "fmt"
	"snp/models"
	//	"strings"

	"github.com/astaxie/beego"
	"time"
	"errors"
	"snp/utils"
)

type PkgController struct {
	beego.Controller
}

func (this *PkgController) Get() {
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	utils.GetLogger().Info("model: pkg, action: %s, id: %s", action, id)

	data := ResObj{}
	switch action {
	case "get":
		if id == "d" {
			id = "default"
		}
		utils.GetConsole().Info("-getting pkg: %s", id)
		data.code = 0
		data.message = ""
		data.data = models.GetPkgFullInfo(id)
		this.Data["json"] = data.Json()
	case "new":
	case "fork":
		utils.GetConsole().Info("forking; %s", id)
		if id == "" {
			this.Data["json"] = false
		} else {
			name := this.GetString("name", "pkg-"+time.Now().Format(time.Stamp))
			pkg, err := forkPkg(id, name)
			if err != nil {
				data.code = -1
				data.message = err.Error()
			} else {
				data.code = 0
				data.message = pkg.Name
			}
		}
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf8")
	this.ServeJSON()
}

/**
 * fork a package
 */
func forkPkg(name string, newName string) (models.Pkg, error){
	// copy package
	old, err := models.GetPkgByName(name)
	if err != nil {
		return models.Pkg{}, err
	}
	// check if newName exist
	chkPkg, _ := models.GetPkgByName(newName)
//	fmt.Println("name:", chkPkg.Name, )
	if chkPkg.Name == newName {
		return models.Pkg{}, errors.New("pkg name already exist")
	}
	return old.Copy(newName)
}

func NewPkg() {

}
