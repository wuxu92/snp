package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
  "snp/utils"
)

type Site struct {
  Id bson.ObjectId  `json:"id" bson:"_id,omitempty"`
  Title string      `json:"title"`
  Url string        `json:"url"`
  Created string    `json:"created"`
  Available bool    `json:"avaliable"`
  HashCode string   `json:"hashcode"`
}

func GetInitSites() []Site {
  createTime := time.Now().Format(time.RFC3339)
//  hasher := md5.New()
  DefaultSites := []Site{
    Site{ bson.NewObjectId(), "facebook", "http://facebook.com", createTime, true, utils.Md5("http://facebook.com")},
    Site{ bson.NewObjectId(), "twitter", "http://twitter.com", createTime, true, utils.Md5("http://twitter.com")},
    Site{ bson.NewObjectId(), "mail.163", "http://mail.163.com/", createTime, true, utils.Md5("http://mail.163.com/")},
    Site{ bson.NewObjectId(), "bitunion", "http://bitunion.org/", createTime, true, utils.Md5("http://bitunion.org/")},
    Site{ bson.NewObjectId(), "baidu", "http://baidu.com", createTime, true, utils.Md5("http://baidu.com")},
    Site{ bson.NewObjectId(), "bitpt", "http://bitpt.cn/", createTime, true, utils.Md5("http://bitpt.cn/")},
    Site{ bson.NewObjectId(), "v2ex", "http://v2ex.com", createTime, true, utils.Md5("http://v2ex.com")},
    Site{ bson.NewObjectId(), "bilibili", "http://www.bilibili.com/", createTime, true, utils.Md5("http://www.bilibili.com/")},
    Site{ bson.NewObjectId(), "zhihu", "http://www.zhihu.com/", createTime, true, utils.Md5("http://www.zhihu.com/")},
    Site{ bson.NewObjectId(), "leanpub.com", "https://leanpub.com/GoNotebook/read", createTime, true, utils.Md5("https://leanpub.com/GoNotebook/read")},
    Site{ bson.NewObjectId(), "ruanyifeng", "http://javascript.ruanyifeng.com/", createTime, true, utils.Md5("http://javascript.ruanyifeng.com/")},
    Site{ bson.NewObjectId(), "laravel.com", "http://laravel.com/docs/5.1/installation", createTime, true, utils.Md5("http://laravel.com/docs/5.1/installation")},
    Site{ bson.NewObjectId(), "10.4.16.95:85", "http://10.4.16.95:85/project/laravel/public/", createTime, true, utils.Md5("http://10.4.16.95:85/project/laravel/public/")},
    Site{ bson.NewObjectId(), "10.4.16.91", "http://10.4.16.91:9090/", createTime, true, utils.Md5("http://10.4.16.91:9090/")},
    Site{ bson.NewObjectId(), "leetcode", "https://leetcode.com/", createTime, true, utils.Md5("https://leetcode.com/")},
    Site{ bson.NewObjectId(), "aliyun.com", "https://ecs.console.aliyun.com/#/home", createTime, true, utils.Md5("https://ecs.console.aliyun.com/#/home")},
    Site{ bson.NewObjectId(), "clctrip", "http://www.clctrip.com/", createTime, true, utils.Md5("http://www.clctrip.com/")},
    Site{ bson.NewObjectId(), "aws.amazon.com", "https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#", createTime, true, utils.Md5("https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#")},
    Site{ bson.NewObjectId(), "bt.neu6.edu.cn", "http://bt.neu6.edu.cn/forum.php", createTime, true, utils.Md5("http://bt.neu6.edu.cn/forum.php")},
    Site{ bson.NewObjectId(), "tower.im", "https://tower.im/teams/ae458404532a46958b30263b7453fd21/projects/", createTime, true, utils.Md5("https://tower.im/teams/ae458404532a46958b30263b7453fd21/projects/")},
    Site{ bson.NewObjectId(), "tower", "http://tower.im", createTime, true, utils.Md5("http://tower.im")},
    Site{ bson.NewObjectId(), "210.76.97.52", "https://210.76.97.52/", createTime, true, utils.Md5("https://210.76.97.52/")},
    Site{ bson.NewObjectId(), "openmymind", "http://openmymind.net/", createTime, true, utils.Md5("http://openmymind.net/")},
    Site{ bson.NewObjectId(), "research.swtch", "http://research.swtch.com/", createTime, true, utils.Md5("http://research.swtch.com/")},
    Site{ bson.NewObjectId(), "golanghome", "http://golanghome.com/", createTime, true, utils.Md5("http://golanghome.com/")},
    Site{ bson.NewObjectId(), "riyu", "http://riyu.io/", createTime, true, utils.Md5("http://riyu.io/")},
    Site{ bson.NewObjectId(), "outlook", "http://outlook.com", createTime, true, utils.Md5("http://outlook.com")},
    Site{ bson.NewObjectId(), "stackedit", "https://stackedit.io/editor", createTime, true, utils.Md5("https://stackedit.io/editor")},
    Site{ bson.NewObjectId(), "shadowsocks", "https://shadowsocks.biz/", createTime, true, utils.Md5("https://shadowsocks.biz/")},
    Site{ bson.NewObjectId(), "cn.mathworks", "http://cn.mathworks.com/help/fuzzy/what-is-sugeno-type-fuzzy-inference.html", createTime, true, utils.Md5("http://cn.mathworks.com/help/fuzzy/what-is-sugeno-type-fuzzy-inference.html")},
    Site{ bson.NewObjectId(), "beej.us", "http://beej.us/guide/bgnet/output/html/multipage/index.html", createTime, true, utils.Md5("http://beej.us/guide/bgnet/output/html/multipage/index.html")},
    Site{ bson.NewObjectId(), "niuxss", "https://niuxss.com/", createTime, true, utils.Md5("https://niuxss.com/")},
    Site{ bson.NewObjectId(), "solidot", "http://www.solidot.org/", createTime, true, utils.Md5("http://www.solidot.org/")},
    Site{ bson.NewObjectId(), "stackoverflow.com/questions/tagged", "http://stackoverflow.com/questions/tagged/c", createTime, true, utils.Md5("http://stackoverflow.com/questions/tagged/c")},
    Site{ bson.NewObjectId(), "c-faq.com/index", "http://c-faq.com/index.html", createTime, true, utils.Md5("http://c-faq.com/index.html")},
    Site{ bson.NewObjectId(), "changelog.com/tagged/go", "https://changelog.com/tagged/go/", createTime, true, utils.Md5("https://changelog.com/tagged/go/")},
    Site{ bson.NewObjectId(), "lwn.net", "https://lwn.net/", createTime, true, utils.Md5("https://lwn.net/")},
    Site{ bson.NewObjectId(), "cloudbus.org", "http://www.cloudbus.org/cloudsim/", createTime, true, utils.Md5("http://www.cloudbus.org/cloudsim/")},
    Site{ bson.NewObjectId(), "ibm.com/developerworks/cn/opensource/os-twitterstorm", "https://www.ibm.com/developerworks/cn/opensource/os-twitterstorm/", createTime, true, utils.Md5("https://www.ibm.com/developerworks/cn/opensource/os-twitterstorm/")},
    Site{ bson.NewObjectId(), "news.ycombinator", "https://news.ycombinator.com/", createTime, true, utils.Md5("https://news.ycombinator.com/")},
    Site{ bson.NewObjectId(), "reddit", "https://www.reddit.com/", createTime, true, utils.Md5("https://www.reddit.com/")},
    Site{ bson.NewObjectId(), "c.learncodethehardway.org", "http://c.learncodethehardway.org/book/", createTime, true, utils.Md5("http://c.learncodethehardway.org/book/")},
    Site{ bson.NewObjectId(), "alpha.wallhaven", "http://alpha.wallhaven.cc/", createTime, true, utils.Md5("http://alpha.wallhaven.cc/")},
    Site{ bson.NewObjectId(), "liaoxuefeng.com/wiki", "http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000", createTime, true, utils.Md5("http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000")},
    Site{ bson.NewObjectId(), "csdn.net/article/2014-01-27/2818282-Spark-Streaming-big", "http://www.csdn.net/article/2014-01-27/2818282-Spark-Streaming-big-data", createTime, true, utils.Md5("http://www.csdn.net/article/2014-01-27/2818282-Spark-Streaming-big-data")},
    Site{ bson.NewObjectId(), "infoq.com/cn/news/2015/02/apache-samza-top", "http://www.infoq.com/cn/news/2015/02/apache-samza-top-project", createTime, true, utils.Md5("http://www.infoq.com/cn/news/2015/02/apache-samza-top-project")},
    Site{ bson.NewObjectId(), "wiki.mbalib.com/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE", "http://wiki.mbalib.com/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", createTime, true, utils.Md5("http://wiki.mbalib.com/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91")},
    Site{ bson.NewObjectId(), "zh.wikipedia.org/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", "https://zh.wikipedia.org/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91", createTime, true, utils.Md5("https://zh.wikipedia.org/wiki/%E6%A8%A1%E7%B3%8A%E9%80%BB%E8%BE%91")},
    Site{ bson.NewObjectId(), "csdn.net", "http://www.csdn.net/article/1970-01-01/2820196", createTime, true, utils.Md5("http://www.csdn.net/article/1970-01-01/2820196")},
    Site{ bson.NewObjectId(), "linux.die.net", "http://linux.die.net/man/", createTime, true, utils.Md5("http://linux.die.net/man/")},
    Site{ bson.NewObjectId(), "wsgzao.github.io/post/hexo", "http://wsgzao.github.io/post/hexo-guide/", createTime, true, utils.Md5("http://wsgzao.github.io/post/hexo-guide/")},
    Site{ bson.NewObjectId(), "cn-stage.vuejs", "http://cn-stage.vuejs.org/", createTime, true, utils.Md5("http://cn-stage.vuejs.org/")},
    Site{ bson.NewObjectId(), "overapi", "http://overapi.com/", createTime, true, utils.Md5("http://overapi.com/")},
    Site{ bson.NewObjectId(), "zhangxinxu", "http://www.zhangxinxu.com/wordpress/", createTime, true, utils.Md5("http://www.zhangxinxu.com/wordpress/")},
    Site{ bson.NewObjectId(), "devtf", "http://www.devtf.cn/", createTime, true, utils.Md5("http://www.devtf.cn/")},
    Site{ bson.NewObjectId(), "blog.zts1993.com/5170", "http://blog.zts1993.com/5170.html", createTime, true, utils.Md5("http://blog.zts1993.com/5170.html")},
    Site{ bson.NewObjectId(), "labix.org/mgo", "https://labix.org/mgo", createTime, true, utils.Md5("https://labix.org/mgo")},
    Site{ bson.NewObjectId(), "developer.mozilla.org/zh-CN/docs/Web/JavaScript/A_re-introduction_to_JavaScript", "https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/A_re-introduction_to_JavaScript", createTime, true, utils.Md5("https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/A_re-introduction_to_JavaScript")},
    Site{ bson.NewObjectId(), "portal.budgetvm.com/customer/dashboard", "https://portal.budgetvm.com/customer/dashboard", createTime, true, utils.Md5("https://portal.budgetvm.com/customer/dashboard")},
    Site{ bson.NewObjectId(), "apps.twitter", "https://apps.twitter.com/", createTime, true, utils.Md5("https://apps.twitter.com/")},
    Site{ bson.NewObjectId(), "github.com/twitter/hbc", "https://github.com/twitter/hbc", createTime, true, utils.Md5("https://github.com/twitter/hbc")},
    Site{ bson.NewObjectId(), "code.tutsplus.com/tutorials/building-with-the-twitter-api-using-real-time-streams--cms", "http://code.tutsplus.com/tutorials/building-with-the-twitter-api-using-real-time-streams--cms-22194", createTime, true, utils.Md5("http://code.tutsplus.com/tutorials/building-with-the-twitter-api-using-real-time-streams--cms-22194")},
  }
  return DefaultSites
}
