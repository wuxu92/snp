package controllers

import (
	"github.com/astaxie/beego"
	"snp/utils"
	"snp/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type SiteController struct {
	beego.Controller
}

func (this *SiteController) Get() {
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	if len(id) == 0 {
		this.Data["json"] = nil
		this.ServeJson()
	}
	switch action {
	case "get":
		this.Data["json"] = getSite(id)
		this.ServeJson()
	case "update":
		updateSite(id)
	case "delete":
		deleteSite(id)
	}
}

func getSite(id string) models.Site {
	c := utils.GetMgc().GetDB().C("site")
	site := models.Site{}
	err := c.FindId(bson.ObjectIdHex(id)).One(&site)
	if err != nil {
		fmt.Println("not site find for: ", id)
	}
	return site
}

func updateSite(id string) bool {

	return true
}

func deleteSite(id string) bool {

	return true
}
