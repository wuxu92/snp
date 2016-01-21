package models

import (
	"errors"
	"snp/utils"

	"gopkg.in/mgo.v2/bson"
	"time"
)

type Pkg struct {
	Id              bson.ObjectId   `json:"-" bson:"_id,omitempty"`
	Created         string          `json:"-"`
	Title           string          `json:"title"`
	Owner           string          `json:"owner"`
	Name            string          `json:"name"`
	Password        string          `json:"-"`
	Public          bool            `json:"public"`
	Groups          []bson.ObjectId `json:"groups"`
	Theme           string          `json:"theme"`
	AdditionalStyle string          `json:"additionalStyle"`
	Version         int             `json:"version"`
}

func (this *Pkg) GetGroups() map[string]Group {
	grpLen := len(this.Groups)
	if grpLen == 0 {
		return nil
	}
	grps := make(map[string]Group, grpLen)
	mgc := utils.GetMgc()
	c := mgc.GetDB().C("grp")
	for _, id := range this.Groups {
		grp := Group{}
		err := c.FindId(id).One(&grp)
		utils.ErrChk(err)

		grps[id.Hex()] = grp
	}

	return grps
}

func (this *Pkg) GetSites(grps map[string]Group) map[string]Site {
	if grps == nil {
		grps = this.GetGroups()
	}
	if len(grps) == 0 {
		return nil
	}
	sites := make(map[string]Site)
	for _, grp := range grps {
		tmpSites := grp.GetSites()
		// merge two maps
		// use for loop for now
		for sId, site := range tmpSites {
			sites[sId] = site
		}
	}
	return sites
}

func (this *Pkg) Copy(newName string) (Pkg, error) {
	pkg := Pkg{
		bson.NewObjectId(),
		time.Now().Format(time.RFC3339),
		this.Title,
		this.Owner,
		newName,
		"123456",
		true,
		nil,
		"",
		"",
		1,
	}
	// copy groups
	newGrps := make([]bson.ObjectId, len(this.Groups))
	for idx, grpIdx := range this.Groups {
		old := GetGroupById(grpIdx)
		tmpGrp := old.Copy()
		newGrps[idx] = tmpGrp.Id
	}
	pkg.Groups  =newGrps

	// save pkg to mongo
	c := utils.GetMgc().GetDB().C("pkg")
	err := c.Insert(&pkg)
	if err != nil {
		return pkg, err
	}
	return pkg, nil
}

func GetPkgFullInfo(name string) map[string]interface{} {
	pkg, _ := GetPkgByName(name)
	data := make(map[string]interface{})
	data["pkg"] = pkg
	groups := pkg.GetGroups()
	data["groups"] = groups

	// sites := make(map[string]models.Site)

	data["sites"] = pkg.GetSites(groups)

	return data
}

func GetPkgByName(name string) (Pkg, error) {
	mgc := utils.GetMgc()
	db := mgc.GetDB()
	c := db.C("pkg")

	var pkg Pkg
	err := c.Find(bson.M{"name": name}).One(&pkg)

	if err != nil {
		return pkg, errors.New("pkg not exist")
	}
	return pkg, nil
}

/**
 * check if a name allready exist
 */
func CheckPkgName(name string) bool {
	mgc := utils.GetMgc()
	c := mgc.GetDB().C("pkg")

	n, err := c.Find(bson.M{"name": name}).Count()
	utils.ErrChk(err)
	if n > 0 {
		return true
	} else {
		return false
	}
}

func GetInitPkg(grps []Group) Pkg {
	pkg := Pkg{
		bson.NewObjectId(),
		"initPkg",
		"Init Package",
		"admin",
		"default",
		"123456",
		true,
		nil,
		"default",
		"",
		1,
	}
	grpCount := len(grps)
	grpIds := make([]bson.ObjectId, grpCount)
	for i := 0; i < grpCount; i++ {
		grpIds[i] = grps[i].Id
	}
	pkg.Groups = grpIds
	return pkg
}
