package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"snp/utils"
)

type Site struct {
	Id        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Title     string        `json:"title"`
	Url       string        `json:"url"`
	Created   string        `json:"-"`
	Available bool          `json:"-"`
	// HashCode  string        `json:"hashcode"`
}

func (this *Site) Copy() Site {
	site := Site {
		bson.NewObjectId(),
		this.Title,
		this.Url,
		time.Now().Format(time.RFC3339),
		true,
	}
	c := utils.GetMgc().GetDB().C("site")
	err := c.Insert(&site)
	utils.ErrChk(err)
	return site
}

func (this *Site) Update() error {
	c := utils.GetMgc().GetDB().C("site")
	return c.UpdateId(this.Id, this)
}

func GetSiteById(id bson.ObjectId) Site {
	c := utils.GetMgc().GetDB().C("site")
	site := Site{}
	err := c.FindId(id).One(&site)
	utils.ErrChk(err)
	return site
}

func IsSiteExist(id bson.ObjectId) bool {
	c := utils.GetMgc().GetDB().C("site")
	n, err := c.FindId(id).Count()
	if err !=  nil {
		return false
	} else {
		return n > 0
	}
}

func GetInitSites() []Site {
	createTime := time.Now().Format(time.RFC3339)
	//  hasher := md5.New()
	DefaultSites := []Site{
		Site{bson.NewObjectId(), "facebook", "http://facebook.com", createTime, true},
		Site{bson.NewObjectId(), "twitter", "http://twitter.com", createTime, true},
		Site{bson.NewObjectId(), "mail.163", "http://mail.163.com/", createTime, true},
		Site{bson.NewObjectId(), "bitunion", "http://bitunion.org/", createTime, true},
		Site{bson.NewObjectId(), "baidu", "http://baidu.com", createTime, true},
		Site{bson.NewObjectId(), "bitpt", "http://bitpt.cn/", createTime, true},
		Site{bson.NewObjectId(), "v2ex", "http://v2ex.com", createTime, true},
		Site{bson.NewObjectId(), "bilibili", "http://www.bilibili.com/", createTime, true},
		Site{bson.NewObjectId(), "zhihu", "http://www.zhihu.com/", createTime, true},
		Site{bson.NewObjectId(), "leanpub.com", "https://leanpub.com/GoNotebook/read", createTime, true},
		Site{bson.NewObjectId(), "ruanyifeng", "http://javascript.ruanyifeng.com/", createTime, true},
		Site{bson.NewObjectId(), "laravel.com", "http://laravel.com/docs/5.1/installation", createTime, true},
		Site{bson.NewObjectId(), "10.4.16.95:85", "http://10.4.16.95:85/project/laravel/public/", createTime, true},
		Site{bson.NewObjectId(), "10.4.16.91", "http://10.4.16.91:9090/", createTime, true},
		Site{bson.NewObjectId(), "leetcode", "https://leetcode.com/", createTime, true},
		Site{bson.NewObjectId(), "aliyun.com", "https://ecs.console.aliyun.com/#/home", createTime, true},
		Site{bson.NewObjectId(), "clctrip", "http://www.clctrip.com/", createTime, true},
		Site{bson.NewObjectId(), "aws.amazon.com", "https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#", createTime, true},
		Site{bson.NewObjectId(), "bt.neu6.edu.cn", "http://bt.neu6.edu.cn/forum.php", createTime, true},
		Site{bson.NewObjectId(), "tower.im", "https://tower.im/teams/ae458404532a46958b30263b7453fd21/projects/", createTime, true},
		Site{bson.NewObjectId(), "tower", "http://tower.im", createTime, true},
		Site{bson.NewObjectId(), "210.76.97.52", "https://210.76.97.52/", createTime, true},
		Site{bson.NewObjectId(), "openmymind", "http://openmymind.net/", createTime, true},
		Site{bson.NewObjectId(), "research.swtch", "http://research.swtch.com/", createTime, true},
		Site{bson.NewObjectId(), "golanghome", "http://golanghome.com/", createTime, true},
		Site{bson.NewObjectId(), "riyu", "http://riyu.io/", createTime, true},
		Site{bson.NewObjectId(), "outlook", "http://outlook.com", createTime, true},
		Site{bson.NewObjectId(), "stackedit", "https://stackedit.io/editor", createTime, true},
		Site{bson.NewObjectId(), "shadowsocks", "https://shadowsocks.biz/", createTime, true},
		Site{bson.NewObjectId(), "cn.mathworks", "http://cn.mathworks.com/help/fuzzy/what-is-sugeno-type-fuzzy-inference.html", createTime, true},
		Site{bson.NewObjectId(), "beej.us", "http://beej.us/guide/bgnet/output/html/multipage/index.html", createTime, true},
		Site{bson.NewObjectId(), "niuxss", "https://niuxss.com/", createTime, true},
		Site{bson.NewObjectId(), "solidot", "http://www.solidot.org/", createTime, true},
		Site{bson.NewObjectId(), "stackoverflow.com/questions/tagged", "http://stackoverflow.com/questions/tagged/c", createTime, true},
		Site{bson.NewObjectId(), "c-faq.com/index", "http://c-faq.com/index.html", createTime, true},
		Site{bson.NewObjectId(), "changelog.com/tagged/go", "https://changelog.com/tagged/go/", createTime, true},
		Site{bson.NewObjectId(), "lwn.net", "https://lwn.net/", createTime, true},
		Site{bson.NewObjectId(), "cloudbus.org", "http://www.cloudbus.org/cloudsim/", createTime, true},
		Site{bson.NewObjectId(), "ibm.com/developerworks/cn/opensource/os-twitterstorm", "https://www.ibm.com/developerworks/cn/opensource/os-twitterstorm/", createTime, true},
		Site{bson.NewObjectId(), "news.ycombinator", "https://news.ycombinator.com/", createTime, true},
		Site{bson.NewObjectId(), "reddit", "https://www.reddit.com/", createTime, true},
		Site{bson.NewObjectId(), "c.learncodethehardway.org", "http://c.learncodethehardway.org/book/", createTime, true},
		Site{bson.NewObjectId(), "alpha.wallhaven", "http://alpha.wallhaven.cc/", createTime, true},
		Site{bson.NewObjectId(), "liaoxuefeng.com/wiki", "http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000", createTime, true},
		Site{bson.NewObjectId(), "csdn.net/article/2014-01-27/2818282-Spark-Streaming-big", "http://www.csdn.net/article/2014-01-27/2818282-Spark-Streaming-big-data", createTime, true},
		Site{bson.NewObjectId(), "infoq.com/cn/news/2015/02/apache-samza-top", "http://www.infoq.com/cn/news/2015/02/apache-samza-top-project", createTime, true},
		Site{bson.NewObjectId(), "wiki.mbalib.com/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE", "http://wiki.mbalib.com/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", createTime, true},
		Site{bson.NewObjectId(), "zh.wikipedia.org/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", "https://zh.wikipedia.org/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", createTime, true},
		Site{bson.NewObjectId(), "csdn.net", "http://www.csdn.net/article/1970-01-01/2820196", createTime, true},
		Site{bson.NewObjectId(), "linux.die.net", "http://linux.die.net/man/", createTime, true},
		Site{bson.NewObjectId(), "wsgzao.github.io/post/hexo", "http://wsgzao.github.io/post/hexo-guide/", createTime, true},
		Site{bson.NewObjectId(), "cn-stage.vuejs", "http://cn-stage.vuejs.org/", createTime, true},
		Site{bson.NewObjectId(), "overapi", "http://overapi.com/", createTime, true},
		Site{bson.NewObjectId(), "zhangxinxu", "http://www.zhangxinxu.com/wordpress/", createTime, true},
		Site{bson.NewObjectId(), "devtf", "http://www.devtf.cn/", createTime, true},
		Site{bson.NewObjectId(), "blog.zts1993.com/5170", "http://blog.zts1993.com/5170.html", createTime, true},
		Site{bson.NewObjectId(), "labix.org/mgo", "https://labix.org/mgo", createTime, true},
		Site{bson.NewObjectId(), "developer.mozilla.org/zh-CN/docs/Web/JavaScript/A_re-introduction_to_JavaScript", "https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/A_re-introduction_to_JavaScript", createTime, true},
		Site{bson.NewObjectId(), "portal.budgetvm.com/customer/dashboard", "https://portal.budgetvm.com/customer/dashboard", createTime, true},
		Site{bson.NewObjectId(), "apps.twitter", "https://apps.twitter.com/", createTime, true},
		Site{bson.NewObjectId(), "github.com/twitter/hbc", "https://github.com/twitter/hbc", createTime, true},
		Site{bson.NewObjectId(), "code.tutsplus.com/tutorials/building-with-the-twitter-api-using-real-time-streams--cms", "http://code.tutsplus.com/tutorials/building-with-the-twitter-api-using-real-time-streams--cms-22194", createTime, true},
	}
	return DefaultSites
}
