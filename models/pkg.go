package models

import (
	"snp/utils"

	"gopkg.in/mgo.v2/bson"
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
	mgc := utils.Mgc{}
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
