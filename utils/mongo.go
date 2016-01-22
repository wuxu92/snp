package utils

import (
	"gopkg.in/mgo.v2"
	//  "gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
	//  "strings"
	// "fmt"
	"time"
)

type Mgc struct {
	session *mgo.Session
}

var instance *Mgc = nil

func GetMgc() *Mgc {
	if instance == nil {
		GetConsole().Info("new mgc client")
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
		GetConsole().Info("connect to mongo: %s", host)
		session, err := mgo.Dial(host)
		ErrChk(err)
		mgc.session = session
		mgc.session.SetMode(mgo.Monotonic, true)
		mgc.session.SetSocketTimeout(60 * time.Second)
	} else {
		// GetConsole().Info("reuse session")
	}
	return mgc.session
}

func (this *Mgc) GetRawSession() *mgo.Session {
	return this.session
}

func (this *Mgc) GetDB() *mgo.Database {
	dbName := beego.AppConfig.String("mgoDB")
	if len(dbName) == 0 {
		dbName = "snp"
	}
	db := this.GetSession().DB(dbName)
	return db
}

//func (this *Mgc) Insert(col string, ins interface{}, others ...interface{}) bool {
//	c := this.GetDB().C(col)
//	c.Insert(ins)
//	return c.Insert(others)
//}

func (this *Mgc) Close() {
	this.session.Close()
}

func GetDBStatic() *mgo.Database {
	mgc := Mgc{}
	return mgc.GetDB()
}
