package models

import (
	"fmt"
	"snp/utils"
	"time"

	"gopkg.in/mgo.v2/bson"
	"reflect"
)

type Group struct {
	Id              bson.ObjectId   `json:"-" bson:"_id,omitempty"`
	Created         string          `json:"-"`
	Title           string          `json:"title"`
	Owner           string          `json:"owner"`
	GroupLink       []string        `json:"-"`
	Followable      bool            `json:"followable"`
	Localized       bool            `json:"localized"`
	FollowFrom      string          `json:"followfrom"`
	AdditionalStyle string          `json:"additionalStyle"`
	Sites           []bson.ObjectId `json:"sites"`
	Version         int             `json:"version"`
}

func (this *Group) GetSites() map[string]Site {
	if len(this.Sites) == 0 {
		return nil
	}
	sites := make(map[string]Site)
	c := utils.GetMgc().GetDB().C("site")
	//  fmt.Println(this.Title, "has sites: ", len(this.Sites))
	for _, sid := range this.Sites {
		site := Site{}
		err := c.FindId(sid).One(&site)
		utils.ErrChk(err)

		sites[sid.Hex()] = site
	}
	return sites
}

func (this *Group) Copy() Group{
	grp := Group{
		bson.NewObjectId(),
		time.Now().Format(time.RFC3339),
		this.Title,
		this.Owner,
		this.GroupLink,
		true,
		true,
		this.Id.Hex(),
		this.AdditionalStyle,
		nil,
		1,
	}
	sites := make([]bson.ObjectId, len(this.Sites))
	for idx, siteId := range this.Sites {
		old := GetSiteById(siteId)
		tmpSite := old.Copy()
		sites[idx] = tmpSite.Id
	}
	grp.Sites = sites
	// save group to mongo
	c := utils.GetMgc().GetDB().C("grp")
	err := c.Insert(&grp)
	utils.ErrChk(err)
	return grp
}

func (this *Group) HasSite(url string) bool {
	for _, s := range this.GetSites() {
		if s.Url == url {
			return true
		}
	}
	return false
}

func (this *Group) HasSiteId(id bson.ObjectId) bool {
	for _, sid := range this.Sites {
		if reflect.DeepEqual(sid, id) {
			return true
		}
	}
	return false
}
func (this *Group) AddSite(title, url string) (Site, error) {
	site := Site{
		bson.NewObjectId(),
		title,
		url,
		time.Now().Format(time.RFC3339),
		true,
	}
	err := utils.GetMgc().GetDB().C("site").Insert(site)
	if err != nil {
		return Site{}, err
	}
	utils.GetLogger().Info("insert site: %s", site.Id.Hex())
	this.Sites = append(this.Sites, site.Id)
	this.Update()
	return site, nil
}

func (this *Group) Update() bool {
	c := utils.GetMgc().GetDB().C("grp")
	err := c.UpdateId(this.Id, this)
	return err == nil
}

func GetGroupById(id bson.ObjectId) Group {
	// db := utils.GetMgc().GetDB()
	c := utils.GetMgc().GetDB().C("grp")
	grp := Group{}
	err := c.FindId(id).One(&grp)
	utils.ErrChk(err)
	return grp
}

func IsGroupExist(id bson.ObjectId) bool {
	c := utils.GetMgc().GetDB().C("grp")
	count, err := c.FindId(id).Count()
	if err != nil {
		return false
	}
	return count > 0
}

func GetGroupFromSites(title string, sites []Site, start, length int) Group {
	group := Group{}
	group.Id = bson.NewObjectId()
	group.Created = time.Now().Format(time.RFC3339)
	group.Title = title
	group.Owner = "wuxu"
	group.GroupLink = nil
	group.Followable = true
	group.Localized = true
	group.FollowFrom = ""
	group.AdditionalStyle = ""
	var siteArr = make([]bson.ObjectId, length)
	var sitesLen = len(sites)
	//  var end = start + length
	if start+length > sitesLen {
		length = sitesLen - start
	}
	fmt.Println("grp:", length, start, len(sites))
	for i := 0; i < length; i++ {
		siteArr[i] = sites[start+i].Id
	}
	group.Sites = siteArr
	group.Version = 1

	return group
}
