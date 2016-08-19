package controllers

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/zituocn/M.VMovie/models"
	"strconv"
	"strings"
	"time"
)

const (
	host = "http://127.0.0.1:8080"
)

type IndexHandle struct {
	baseController
}

//新剧推荐
func (this *IndexHandle) New() {
	this.Ctx.Output.Header("Cache-Control", "public")
	url := host + "/api/new"
	info := GetNewList(url)

	this.Data["list"] = info.List
	this.TplName = "_new.html"
}

//资讯详情
func (this *IndexHandle) Page() {

	///get param
	idstr := this.Ctx.Input.Param(":id")
	this.Ctx.Output.Header("Cache-Control", "public")
	//request json
	url := host + "/api/article/" + idstr + "/"
	info := GetPageInfo(url)
	this.Data["info"] = info.Info
	this.TplName = "_pageinfo.html"
}

///资讯列表
func (this *IndexHandle) News() {
	var (
		pager string
	)
	this.Ctx.Output.Header("Cache-Control", "public")
	///get param
	pagestr := this.Ctx.Input.Param(":page")
	page, _ := strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}

	//request json
	url := host + "/api/news/" + pagestr + "/"
	info := GetNewsList(url)

	pager = this.PageList(10, page, info.RecordCount, false, "/news")
	this.Data["pager"] = pager
	this.Data["list"] = info.List

	this.TplName = "_news.html"
}

//今日更新的影片
func (this *IndexHandle) Today() {

	this.Ctx.Output.Header("Cache-Control", "public")
	url := host + "/api/today"
	info := GetToday(url)

	week := this.GetWeekString()
	this.Data["week"] = week
	this.Data["list"] = info.List
	this.TplName = "_today.html"
}

//搜索页面
func (this *IndexHandle) Search() {
	var (
		pager   string
		keyword string
	)
	this.Ctx.Output.Header("Cache-Control", "public")
	///get param
	keyword = strings.TrimSpace(this.Ctx.Input.Param(":key"))
	pagestr := this.Ctx.Input.Param(":page")
	page, _ := strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}

	//request json
	url := host + "/api/search/" + keyword + "/" + pagestr + "/"
	info := GetSearch(url)

	pager = this.PageList(10, page, info.RecordCount, false, "/search/"+keyword)
	this.Data["pager"] = pager
	this.Data["list"] = info.List
	this.TplName = "_search.html"
}

//详情页面输出
func (this *IndexHandle) Detail() {
	var (
		isend    string
		downitem string
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	id := this.Ctx.Input.Param(":id")
	url := host + "/api/v/" + id
	info := GetDetail(url)

	if info.Minfo == nil {
		this.Abort("404")
	}

	///相关影片
	var item string
	liststring := []string{}
	if len(info.SameList) > 0 {
		item = "<p>相关影片：<br />"
		for _, i := range info.SameList {
			item = item + fmt.Sprintf("<a href=\"/v/%d/\">%s</a><br />", i.Id, i.Name)
		}
		item = item + "</p>"
		liststring = append(liststring, item)
	}

	//下载地址
	count := int64(len(info.DownList))
	for i := 1; int64(i) < (count + 1); i++ {
		hdurl := info.DownList[i-1].Hdtvurl
		if strings.Contains(hdurl, "mkv") {
			if i < 10 {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\"><a href=\"%s\">第0%d集.HDTV.1024.中文字幕.mkv</a></li>", i, hdurl, i)
			} else {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\"><a href=\"%s\">第%d集.HDTV.1024.中文字幕.mkv</a></li>", i, hdurl, i)
			}
		} else {
			if i < 10 {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\"><a href=\"%s\">第0%d集.HDTV.1024.中文字幕.mp4</a></li>", i, hdurl, i)
			} else {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\"><a href=\"%s\">第%d集.HDTV.1024.中文字幕.mp4</a></li>", i, hdurl, i)
			}
		}
	}
	if count < (info.Minfo.Episode + 1) {
		for i := (count + 1); int64(i) < (info.Minfo.Episode + 1); i++ {
			if i < 10 {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\">第0%d集.HDTV.1024.中文字幕.mp4</li>", i, i)
			} else {
				downitem = downitem + fmt.Sprintf("<li id=\"hdtv%d\">第%d集.HDTV.1024.中文字幕.mp4</li>", i, i)
			}
		}
	}

	info.Minfo.Content = strings.Replace(info.Minfo.Content, "\r\n\r\n", "\r\n", -1)
	info.Minfo.Content = strings.Replace(info.Minfo.Content, "\r\n", "<br />", -1)
	if info.Minfo.Isend == 1 {
		isend = "已完结"
	} else {
		isend = fmt.Sprintf("每周%d更新", info.Minfo.Updateweek)
	}
	this.Data["isend"] = isend
	this.Data["rmlist"] = strings.Join(liststring, "\n") //相关影片输出
	this.Data["downitem"] = downitem                     //下载地址输出
	this.Data["info"] = info.Minfo
	this.Data["cinfo"] = info.Cinfo
	this.TplName = "_detial.html"
}

//首页页面输出
func (this *IndexHandle) Index() {
	var (
		week string
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	url := host + "/api/"
	info := GetIndex(url)

	wlist := info.List[0].MList //今日影片
	hlist := info.List[1].MList //推荐影片

	week = this.GetWeekString()
	this.Data["week"] = week
	this.Data["wlist"] = wlist
	this.Data["hlist"] = hlist

	this.TplName = "_index.html"
}

//列表页输出
func (this *IndexHandle) List() {
	var (
		pager string
	)
	///get param
	cid := this.Ctx.Input.Param(":cid")
	pagestr := this.Ctx.Input.Param(":page")
	page, _ := strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}

	//request json
	url := host + "/api/m/" + cid + "/" + pagestr + "/"
	info := GetList(url)

	if info.Cinfo == nil || len(info.Cinfo.Name) == 0 {
		this.Abort("404")
	}

	pager = this.PageList(10, page, info.RecordCount, false, "/m/"+cid)
	this.Data["pager"] = pager
	this.Data["list"] = info.MList
	this.Data["cid"] = cid
	this.Data["cinfo"] = info.Cinfo

	this.TplName = "_list.html"
}

///获取api数据

//获取新剧推荐列表
func GetNewList(url string) models.ApiNewInfo {
	var info models.ApiNewInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

//获取资讯列表数据
func GetNewsList(url string) models.ApiPageListInfo {
	var info models.ApiPageListInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

//获取资讯详情
func GetPageInfo(url string) models.ApiPageDetailInfo {
	var info models.ApiPageDetailInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

//首页数据
func GetIndex(url string) models.ApiIndexInfo {
	var info models.ApiIndexInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

///今日更新的片源
func GetToday(url string) models.ApiTodayInfo {
	var info models.ApiTodayInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

//获取远程列表页数据
func GetList(url string) models.ApiListInfo {
	var list models.ApiListInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&list)
	return list
}

//获取远程详情页面的数据
func GetDetail(url string) models.ApiDetailInfo {
	var info models.ApiDetailInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}

//获取搜索页面的数据
func GetSearch(url string) models.ApiSearchInfo {
	var info models.ApiSearchInfo
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.ToJSON(&info)
	return info
}
