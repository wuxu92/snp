package controllers

import (
	"github.com/astaxie/beego"
	"snp/models"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"snp/utils"
)

type NewModelController struct {
	beego.Controller
}

func (this *NewModelController) Get() {

}

func (this *NewModelController) Post() {
	model := this.Ctx.Input.Param(":model")

	data := ResObj{}
	switch model {
	case "site" :
		title := this.GetString("title")
		url := this.GetString("url")
		grp := this.GetString("grp")

		utils.GetLogger().Info("add site %s - %s to %s", title, url, grp)

		site, err := newSite(title, url, grp)
		if err == nil {
			data.code = 0

			// config data
			reSite := make(map[string]interface{})
			reSite["title"] = site.Title
			reSite["url"]   = site.Url
			reSite["id"]	  = site.Id.Hex()
			data.data = reSite
		} else {
			data.code = 1
			data.message = err.Error()
		}
		this.Data["json"] = data.Json()
	case "grp" :
	case "pkg" :
	}

	this.ServeJson()
}

func newSite(title, url, grpId string) (models.Site, error) {
	grp := models.GetGroupById(bson.ObjectIdHex(grpId))
//	if grp.Id == nil {
//		return models.Site{}, errors.New("grp not exist")
//	}
	if grp.HasSite(url) {
		return models.Site{}, errors.New("grp already has url")
	}
	return grp.AddSite(title, url)
	// eturn models.Site{}, nil
}