/*
* @Author: wuxu92
* @Date:   2016-01-14 21:29:40
* @Last Modified by:   wuxu92
* @Last Modified time: 2016-01-24 14:51:25
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
    editSite: {
      name: "",
      url: "",
      id: "",
      gid: "",
      sIdx: "",
      gIdx: ""
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
      var site = {
        name: this.newSite.name,
        url: this.newSite.url,
        id: ""
      }
      var grpIndx = this.addSiteIndex
      // ajax request add site
      var uri = addSiteUri
      $.ajax({
        url: uri,
        type: 'post',
        dataType: 'json',
        data: {
          grp: this.groups[grpIndx].id,
          title: site.name,
          url: site.url
        }
      })
      .done(function(data) {
        if (typeof data == "string")
          data = $.parseJSON(data)
        if (data.code != 0) {
          console.log("add site error: " + data.message)
          alert(data.message)
          return
        }
        site.id = data.data.id
        grpsVM.groups[grpIndx].sites.push(site)
        console.log("success");
        grpsVM.newSite.name = ""
        grpsVM.newSite.url = ""
      })
      .fail(function() {
        console.log("error");
      })
      .always(function() {
        console.log("complete");
      });
      $("#new-site-modal").modal('hide')
    },
    editGroupSites: function(gIdx, event) {
      // change this group's color
      var nodes = $("a.site[data-gidx="+gIdx+"]")
      var curNode = $(event.target)
      // console.log(nodes)
      if (curNode.hasClass("editing-site")) {
        nodes.removeClass('editSite')
        curNode.removeClass("editing-site")
      } else {
        nodes.addClass('editSite')
        curNode.addClass("editing-site")
      }

    },
    editSiteModal: function(siteId, event) {
      var node = $(event.target)
      // console.log($(event.target))
      if (!node.hasClass('editSite'))
        return true
      else {
        event.preventDefault()
        event.stopPropagation()
        // return false
      }
      var gIdx = node.data("gidx")
      var sIdx = node.data("sidx")
      console.log("edit:" + gIdx + ", sid:" +sIdx)
      if (typeof sIdx === "undefined" || typeof gIdx === "undefined") {
        return
      }
      // get old site
      var group = this.groups[gIdx]
      if (typeof group === "undefined") {
        console.log("old site not exist: " + gIdx + "/" + sIdx)
        return false
      }
      var old = group.sites[sIdx]
      console.log(group)
      console.log(old)
      this.editSite.name = old.name
      this.editSite.url = old.url
      this.editSite.id = old.id
      this.editSite.gid = group.id
      this.editSite.sIdx = sIdx
      this.editSite.gIdx = gIdx
      // var parent = node.parent('div')
      // node.parent('div').width(parent.width() * 3)
      $('#edit-site-modal').modal('show')
      $('#edit-site-modal input').first().focus()
      event.preventDefault()
      event.stopPropagation()
      return false
    },
    updateSite: function(event) {
      console.log(event)
      var node = $(event.target)
      var site = this.editSite
      if (!site.gid || !site.id) {
        console.log("data err")
        return
      }
      // ajax update site
      $.ajax({
        url: editSiteUri+site.id,
        type: 'POST',
        dataType: 'json',
        data: {
          grp: site.gid,
          title: site.name,
          url: site.url
        },
      })
      .done(function(data) {
        console.log(data)
        if (typeof data == "string")
          data = $.parseJSON(data)
        if (data.code != 0) {
          console.log("update failed:" + data.message)
          alert("error:" + data.message)
          return false
        }
        var originSite = grpsVM.groups[site.gIdx].sites[site.sIdx]
        originSite.name = site.name
        originSite.url = site.url
        for (var i in grpsVM.editSite) {
          grpsVM.editSite[i] = ""
        }
        $('#edit-site-modal').modal('hide')
        // update local show
      })
      .fail(function() {
        console.log("error");
      })
      .always(function() {
        console.log("complete");
      });
    },
    deleteSite: function(event) {
      var site = this.editSite
      if (!site.gid || !site.id) {
        console.log("data err")
        return
      }
      // add form data to uri for delete request
      $.ajax({
        url: deleteSiteUri + site.id +"?grp=" + site.gid,
        type: 'DELETE',
        dataType: 'json',
      })
      .done(function(data) {
        console.log(data)
        if (typeof data === "string") {
          data = $.parseJSON(data)
        }
        if (data.code != 0) {
          alert("删除失败：" + data.message)
          return
        }
        // remove the data from page
        var originSite = grpsVM.groups[site.gIdx].sites[site.sIdx]
        // var grp = grpsVM.groups[site.gIdx].sites
        //var siteIdx = site.sIdx
        grpsVM.groups[site.gIdx].sites.splice(site.sIdx, 1)
        // hide the modal
        $('#edit-site-modal').modal('hide')
        console.log("success");
      })
      .fail(function() {
        console.log("error");
      })
      .always(function() {
        console.log("complete");
      });

    }
  }
})
var pkgUri = "api/pkg/get/default"
var addSiteUri = "api/new/site" // post
var editSiteUri = "api/site/edit/" // +sid?grp=
var deleteSiteUri = "/api/site/delete/" // +sid?grp  delete request
$.ajax({
  url: pkgUri,
  type: 'GET',
  dataType: 'json',
})
.done(function(data) {
  console.log(data);
  if (typeof data == "string") {
    data = $.parseJSON(data);
  }
  data = data.data
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
    tmp.id = idx
    tmp.sites = []
    for (var sIdx in grp.sites) {
      sIdx = grp.sites[sIdx];
      site = data.sites[sIdx]
      tmp.sites.push({
        name: site.title.substr(0, 8),
        url: site.url,
        id: sIdx
      })
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
