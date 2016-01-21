package utils

import (
	"gopkg.in/mgo.v2"
	//  "gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
	//  "strings"
)

type Mgc struct {
	session *mgo.Session
}

var instance *Mgc = nil

func GetMgc() *Mgc {
	if instance == nil {
		instance = new(Mgc)
	}
	return instance
}

func (mgc *Mgc) GetSession() *mgo.Session {
	if mgc.session == nil {
		host := beego.AppConfig.String("mgoHost")
		port := beego.AppConfig.String("port")
		user := beego.AppConfig.String("mgoUser")
		dbName := beego.AppConfig.String("mgoDB")

		password := beego.AppConfig.String("mgoPassword")
		if len(host) == 0 {
			host = "127.0.0.1"
		}
		if len(port) > 0 {
			host += ":" + port
		}
		if len(dbName) == 0 {
			dbName = "snp"
		}
		host += "/" + dbName

		if len(user) > 0 {
			host = "mongodb://" + user + ":" + password + "@" + host
		}
		session, err := mgo.Dial(host)
		ErrChk(err)
		mgc.session = session
		mgc.session.SetMode(mgo.Monotonic, true)
	}
	return mgc.session
}

func (this *Mgc) GetDB() *mgo.Database {
	dbName := beego.AppConfig.String("mgoDB")
	if len(dbName) == 0 {
		dbName = "snp"
	}
	db := this.GetSession().DB(dbName)
	return db
}

func (this *Mgc) Close() {
	this.session.Close()
}

func GetDBStatic() *mgo.Database {
	mgc := Mgc{}
	return mgc.GetDB()
}
