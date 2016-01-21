package controllers

import (
	"fmt"
	"snp/models"
	//	"snp/utils"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type GrpController struct {
	beego.Controller
}

func (this *GrpController) Get() {
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	if len(id) == 0 {
		this.Data["json"] = nil
		this.ServeJson()
	}
	switch action {
	case "get":
		this.Data["json"] = getGroup(id)
		this.ServeJson()
	case "update":
		updateGroup(id)
	case "delete":
		deleteGroup(id)
	}
}

func getGroup(id string) map[string]interface{} {
	fmt.Println("get group", id)
	grp := models.GetGroupById(bson.ObjectIdHex(id))
	data := make(map[string]interface{})
	data["grp"] = grp
	data["sites"] = grp.GetSites()
	return data
}

func updateGroup(name string) bool {

	return true
}

func deleteGroup(name string) bool {

	return true
}
