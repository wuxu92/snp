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

// Post for edit site, request method post
// /api/site/edit/siteId?grp=grpid
func (this *SiteController) Post() {
	action := this.Ctx.Input.Param(":action")
	sid := this.Ctx.Input.Param(":id")
	gid := this.GetString("grp", "")
	title := this.GetString("title", "")
	url := this.GetString("url","")

	utils.GetConsole().Info("update site, action：%s, id: %s, grp:%s", action, sid, gid)

	data := ResObj{}
	if gid == "" || title =="" || url == "" {
		data.code = 1
		data.message = "one of grp/title/url is null"
		this.Data["json"] = data.Json()
		this.ServeJson()
		return
	}
	// get group
	siteId := bson.ObjectIdHex(sid)
	grpId := bson.ObjectIdHex(gid)
	if ge, se := !models.IsGroupExist(grpId), !models.IsSiteExist(siteId); se || ge{
		data.code = 2
		if !se {
			data.message = "site not exist"
		} else if !ge {
			data.message = "group not exist"
		} else {
			data.message = "group and site not exist"
		}
		this.Data["json"] = data.Json()
		this.ServeJson()
		return
	}
	grp := models.GetGroupById(grpId)
	if !grp.HasSiteId(siteId) {
		data.code = 3
		data.message = "grp do't has this site"
		this.Data["json"] = data.Json()
		this.ServeJson()
		return
	}
	site := models.GetSiteById(siteId)
	site.Title = title
	site.Url = url
	err := site.Update()
	if err != nil {
		data.code = 4
		data.message = "update site error: " + err.Error()
	} else {
		data.code = 0
		data.message = "ok"
	}
	this.Data["json"] = data.Json()
	this.ServeJson()
}

// Delete uri: /api/site/delete/:id?grp=gid&p=password delete
// @todo add grp and password
func (this *SiteController) Delete() {
  sid := this.Ctx.Input.Param(":id")
  action := this.Ctx.Input.Param(":action")
  gid := this.GetString("grp")

  data := ResObj{}
  utils.GetConsole().Info("a: %s, site: %s, grp: %s -- %s", action, sid, gid)
  if !bson.IsObjectIdHex(sid) || !bson.IsObjectIdHex(gid) {
    data.code = 1
    data.message = "one of site or grp is illeagal"
    this.Data["json"] = data.Json()
    this.ServeJson()
    return
  }
  siteId := bson.ObjectIdHex(sid)
  grpId := bson.ObjectIdHex(gid)

  if ge, se := !models.IsGroupExist(grpId), !models.IsSiteExist(siteId); se || ge{
    data.code = 2
    if !se {
      data.message = "site not exist"
    } else if !ge {
      data.message = "group not exist"
    } else {
      data.message = "group and site not exist"
    }
    this.json(data)
    return
  }

  grp := models.GetGroupById(grpId)
  if !grp.HasSiteId(siteId) {
    //data.code = 3
    //data.message = "grp don't has this site"
    this.json(ResObj{3, "grp don't has this site", nil})
    return
  }

  // remove site from grp
  grp.RemoveSite(siteId)
  // save to db
  grp.Update()
  this.json(ResObj{0, "ok", nil})
}

func (this *SiteController) json(d ResObj) {
  this.Data["json"] = d.Json()
  this.ServeJson()
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
