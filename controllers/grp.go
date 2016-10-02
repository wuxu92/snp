package controllers

import (
	"fmt"
	"snp/models"
	//	"snp/utils"

	"gopkg.in/mgo.v2/bson"
  "snp/utils"
)

type GrpController struct {
	BaseController
}

func (this *GrpController) Get() {
	action := this.Ctx.Input.Param(":action")
	id := this.Ctx.Input.Param(":id")
	if len(id) == 0 {
		this.Data["json"] = nil
		this.ServeJSON()
	}
	switch action {
	case "get":
		this.Data["json"] = getGroup(id)
		this.ServeJSON()
	case "update":
		updateGroup(id)
	case "delete":
		deleteGroup(id)
	}
}

// delete group uri : /api/grp/delete/:id?p=password&pkg=pkgName   with delete request
// todo check password
func (this *GrpController) Delete() {
  gid := this.Ctx.Input.Param(":id")
  action := this.Ctx.Input.Param(":action")
  password := this.GetString("p")
  pkgName := this.GetString("pkg")
  data := ResObj{}

  // log request
  utils.GetConsole().Info("a: %s, grp: %s, p: %s", action, gid, password)

  pkg, err := models.GetPkgByName(pkgName)
  if err != nil {
    data.SetCode(CODE_NO_SUCH_PKG)
    this.Json(data)
    return
  }

  if !bson.IsObjectIdHex(gid) {
    data.SetCode(CODE_MGO_BAD_ID)
    this.Json(data)
    return
  }
  // todo check password
  if !pkg.CheckPassword(password) {
    data.SetCode(CODE_PASSWORD_ERR)
    this.Json(data)
    return
  }

  // do delete process
  // just remove id from pkg's grp list
  deleted := pkg.RemoveGroup(gid)
  if !deleted {
    data.SetCode(CODE_NO_SUCH_GRP)
  } else {
    data.SetCode(CODE_OK)
  }

  this.Json(data)
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
