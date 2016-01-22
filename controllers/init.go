package controllers

import(
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "snp/models"
  "fmt"
  "snp/utils"
)

type InitController struct {
  beego.Controller
}

func(ctl *InitController) Get() {
  session := utils.GetMgc().GetSession()

  // defer session.Close()
  
  session.SetMode(mgo.Monotonic, true)
  
  c := session.DB("snp").C("site")
  
  // result := make(Site, )
  c.RemoveAll(nil);
  sites := models.GetInitSites()
  siteArr := make([]interface{}, len(sites))
  for idx, _ := range sites {
    siteArr[idx] = sites[idx]
  }
  fmt.Println("to insert site count: ", len(sites))
  err := c.Insert(siteArr...)
  ErrorChk(err)
//  query := c.Find(bson.M{})// .All(&result)
  fmt.Print("insert sites:")
  PrintQuery(c.Find(bson.M{}))
  siteCount, _ := c.Find(bson.M{}).Count()
  
  // create init groups
  grp1 := models.GetGroupFromSites("Tech", sites, 0, 10)
  grp2 := models.GetGroupFromSites("Read", sites, 10, 5)
  grp3 := models.GetGroupFromSites("Mine", sites, 15, 20)
  grp4 := models.GetGroupFromSites("Study", sites, 35, 2)
  grp5 := models.GetGroupFromSites("Music", sites, 37, 6)
  grp6 := models.GetGroupFromSites("Paper", sites, 43, 16)
  cg := session.DB("snp").C("grp")
  cg.RemoveAll(nil)
  err = cg.Insert(grp1, grp2, grp3, grp4, grp5, grp6)
  ErrorChk(err)
  
  fmt.Print("insert groups ")
  PrintQuery(cg.Find(bson.M{}))
  grpCount, _ := cg.Find(bson.M{}).Count()
  
  // int package
  cp := session.DB("snp").C("pkg")
  cp.RemoveAll(nil)
  err = cp.Insert(models.GetInitPkg([]models.Group{grp1, grp2, grp3, grp4, grp5, grp6}))
  ErrorChk(err)
  
  pkgCount, _ := cp.Find(bson.M{}).Count()
  
  fmt.Print("insert package ")
  PrintQuery(cp.Find(bson.M{}))
  
  ctl.Data["grpcount"] = grpCount
  ctl.Data["sitecount"] = siteCount
  ctl.Data["pkgcount"] = pkgCount
  // ctl.Data["Sites"] = &result
  ctl.Data["Website"] = "wuxu92.com"
  ctl.Data["Email"] = "admin@wuxu92.com"
  ctl.Layout = "layout/main.html"
  ctl.TplNames = "init.tpl"
  ctl.Render()
}

func PrintQuery(query *mgo.Query) {
  count, _ := query.Count()
  fmt.Println("query count:", count)
}

func ErrorChk(err error) {
  if err != nil {
    panic(err)
  }
}