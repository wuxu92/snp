package controllers

import (
	"errors"
	"fmt"
	"snp/models"
	"snp/utils"
	"strings"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ApiController struct {
	beego.Controller
}

func (ctl *ApiController) Get() {
	model := ctl.Ctx.Input.Param(":model")
	action := ctl.Ctx.Input.Param(":action")
	id := ctl.Ctx.Input.Param(":id")
	fmt.Println("model:", model, " action: ", action, "id: ", id)

	if strings.Compare(model, "pkg") == 0 &&
		strings.Compare(action, "get") == 0 {

		if id == "" {
			id = "default"
		}
		fmt.Println("getting pkg:", id)
		ctl.Data["json"] = models.GetPkgFullInfo(id)
		ctl.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf8")
		ctl.ServeJSON()
	} else {
		ctl.ServeJSON()
	}
}

func getPkg(name string) (models.Pkg, error) {
	mgc := utils.Mgc{}
	db := mgc.GetDB()
	c := db.C("pkg")

	var pkg models.Pkg
	err := c.Find(bson.M{"name": name}).One(&pkg)

	if err != nil {
		return pkg, errors.New("pkg not exist")
	}
	return pkg, nil
}
