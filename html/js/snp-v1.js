/*
* @Author: wuxu92
* @Date:   2016-01-14 21:29:40
* @Last Modified by:   wuxu92
* @Last Modified time: 2016-01-20 22:19:50
*/

'use strict';
var dfl_sites = {
  name: "default navigation site",
  creator: "xxxx"
}

var dfl_groups = [
  {
    title: "科技",
    name: "tech",
    sites: [
      {
        name: "瘾科技",
        url: "http://cn.engadget.com/"
      },
      {
        name: "36 KR",
        url: "http://36kr.com"
      }
    ]
  },
  {
    title: "社交",
    name: "social network",
    sites: [
      {
        name: "微博",
        url: "http://weibo.com"
      },
      {
        name: "贴吧",
        url: "http://tieba.baidu.com"
      }
    ]
  },
  {
    title: "搜索",
    name: "Search",
    sites: [
      {
        name: "Baidu",
        url: "http://baidu.com"
      },
      {

        name: "Bing",
        url: "http://cn.bing.com"
      },
      {
        name: "Google",
        url: "http://google.ca"
      },
      {
        name: "搜狗",
        url: "https://www.sogou.com/"
      }
    ]
  }
]

var grpsVM = new Vue({
  el: "#navi-container",
  data: {
    groups: dfl_groups,
    addSiteIndex: 0,
    newSite: {
      name: "",
      url: "http://",
    },
    newSiteToGroup: ""
  },
  methods: {
    addSiteToGroup: function(gIdx, event) {
      this.addSiteIndex = gIdx
      this.newSiteToGroup = this.groups[gIdx].title
      $("#new-site-modal").modal('show')
      console.log("add site to group: " + gIdx)
    },
    saveSite: function() {
      this.groups[this.addSiteIndex].sites.push({name: this.newSite.name, url: this.newSite.url})
      this.newSite.name = ""
      this.newSite.url = "http://"
      $("#new-site-modal").modal('hide')
    }
  }
})
var pkgUri = "api/pkg/get/default"
$.ajax({
  url: pkgUri,
  type: 'GET',
  dataType: 'json',
})
.done(function(data) {
  console.log(data);
  if (data instanceof String) {
    data = $.parseJSON(data);
  }
  // build groups
  if (!data.hasOwnProperty("groups") || !data.hasOwnProperty("sites")) {
    console.log("data corrupted")
    return
  }
  var grp, site
  var newGrps = []
  var reGrps = data.groups
  for (var idx in reGrps) {
    grp = reGrps[idx]
    var tmp = {}
    tmp.title = grp.title
    tmp.sites = []
    for (var sIdx in grp.sites) {
      sIdx = grp.sites[sIdx];
      site = data.sites[sIdx]
      tmp.sites.push({name: site.title.substr(0, 8), url: site.url})
    }
    newGrps.push(tmp)
  }
  grpsVM.groups = newGrps
})
.fail(function() {
  console.log("error");
})
.always(function() {
  console.log("complete");
});
