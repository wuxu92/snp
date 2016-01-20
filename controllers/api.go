package controllers

import (
  "github.com/astaxie/beego"
  _ "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "fmt"
  "strings"
  "snp/models"
  "snp/utils"
  "errors"
)

type ApiController struct {
  beego.Controller
}

func (ctl *ApiController) Get() {
  model := ctl.Ctx.Input.Param(":model")
  action := ctl.Ctx.Input.Param(":action")
  id := ctl.Ctx.Input.Param(":id")
  fmt.Println("model:", model, " action: ", action, "id: ", id)
  
  if strings.Compare(model, "pkg")==0 &&
    strings.Compare(action, "get") == 0 {
    // id := ctl.Input().Get("id")
    if id == "" {
      id = "default"
    }
    fmt.Println("getting pkg:", id)
    var pkg models.Pkg
    
    pkg,_ = getPkg(id)
    
    data := make(map[string]interface{})
    data["pkg"] = pkg
    groups := pkg.GetGroups();
    data["groups"] = groups
    
    // sites := make(map[string]models.Site)

    data["sites"] = pkg.GetSites(groups)
    ctl.Data["json"] = &data //&{"pkg": pkg, "grps": "grps", "sites":"sites"}
    ctl.ServeJson()
  } else {
    ctl.ServeJson()
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