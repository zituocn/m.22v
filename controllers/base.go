package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	//"strings"
	"time"
)

type baseController struct {
	beego.Controller
}

func Error(err error) {
	if err != nil {
		panic(err)
		beego.Error(err.Error())
		//os.Exit(1)
	}
}

//返回星期四格式
func (this *baseController) GetWeekString() string {
	var wstring string
	switch time.Now().Weekday() {
	case time.Monday:
		wstring = "一"
	case time.Tuesday:
		wstring = "二"
	case time.Wednesday:
		wstring = "三"
	case time.Thursday:
		wstring = "四"
	case time.Friday:
		wstring = "五"
	case time.Saturday:
		wstring = "六"
	case time.Sunday:
		wstring = "天"
	default:
		wstring = "X"
	}
	return wstring
}

//显示分页链接
func (this *baseController) PageList(pagesize, page, recordcount int64, first bool, path string) (pager string) {
	if recordcount == 0 {
		return ""
	}

	var pagecount int64
	pagecount = 0

	if recordcount%pagesize == 0 {
		pagecount = recordcount / pagesize
	} else {
		pagecount = (recordcount / pagesize) + 1
	}

	//pager = "<span>" + strconv.FormatInt(page, 10) + "/" + strconv.FormatInt(pagecount, 10) + "</span>"

	//if pagecount < 2 {
	//	return "<span>共1页</span>"
	//}

	//pager = pager + "<a href=\"" + path + "/\">第一页</a>"
	if page > 1 {
		if page == 2 {
			pager = pager + "<li class=\"lastpage\"><a href=\"" + path + "/\">上一页</a></li>"
		} else {
			pager = pager + "<li class=\"lastpage\"><a href=\"" + path + "/" + strconv.FormatInt(page-1, 10) + "/\" >上一页</a></li>"
		}
	} else {
		pager = pager + "<li class=\"lastpage\"><a href=\"" + path + "/\">上一页</a></li>"
	}
	//<li class="nextpage"><a href="item_2_3.htm">下一页 ></a></li>

	if page < pagecount {
		pager = pager + "<li class=\"nextpage\"><a href=\"" + path + "/" + strconv.FormatInt(page+1, 10) + "/\">下一页</a></li>"
	} else {
		pager = pager + "<li class=\"nextpage\"><a href=\"" + path + "/" + strconv.FormatInt(pagecount, 10) + "/\">下一页</a></li>"
	}

	//pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(pagecount, 10) + "/\"  class=\"next\">最后一页</a>"

	//pager = pager + "<span >每页" + strconv.FormatInt(pagesize, 10) + "/共" + strconv.FormatInt(recordcount, 10) + "</span>"

	return pager

}
