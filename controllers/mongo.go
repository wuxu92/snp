package controllers
//
//import (
//  "github.com/astaxie/beego"
//  "gopkg.in/mgo.v2"
//  "gopkg.in/mgo.v2/bson"
//  "fmt"
//)
//
//type Site struct {
//  Name string
//  Url string
//  ClickCount int
//}
//
//type MongoController struct{
//  beego.Controller
//}
//
//func (ctl *MongoController) Get(){
//  session, err := mgo.Dial("127.0.0.1:27017")
//  if err != nil {
//    panic(err)
//  }
//  defer session.Close()
//
//  session.SetMode(mgo.Monotonic, true)
//
//  c := session.DB("snp").C("site")
//
//
//  if err != nil {
//    panic(err)
//  }
//  // result := make(Site, )
//  c.RemoveAll(nil);
//  err = c.Insert(&Site{"Baidu", "www.baidu.com", 1},
//    &Site{"v2ex", "v2ex.com", 22})
//  query := c.Find(bson.M{"name":"Baidu"})// .All(&result)
//  count, _ := query.Count()
//  result := make([]Site, count)
//  err = query.All(&result)
//
//  if err != nil {
//    panic(err)
//  }
//  fmt.Println(result)
//  ctl.Data["Index"] = count
//  ctl.Data["Sites"] = &result
//  ctl.Data["Website"] = "wuxu92.com"
//  ctl.Data["Email"] = "admin@wuxu92.com"
//  ctl.Layout = "layout/main.html"
//  ctl.TplNames = "mongo.tpl"
//}