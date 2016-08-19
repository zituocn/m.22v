package models

import (
	"fmt"
)

///详情页输出结构
type ApiDetailInfo struct {
	Minfo    *MovieInfo      //影片信息
	SameList []*MovieInfo    //类似影片
	DownList []*DownAddrInfo //下载地址列表
	Cinfo    *MovieClassInfo //分类信息
}

//列表页输出模型
type ApiListInfo struct {
	Cinfo       *MovieClassInfo //分类信息
	MList       []*MovieInfo    //影片列表
	Page        int64           //当前页码
	RecordCount int64           //记当总数
}

type ApiIndexMovieList struct {
	MList []*MovieInfo //影片小列表
}

//首页输出模型
type ApiIndexInfo struct {
	List []*ApiIndexMovieList
}

//今日更新列表
type ApiTodayInfo struct {
	List []*MovieInfo //影片列表
}

//搜索模型
type ApiSearchInfo struct {
	List        []*MovieInfo //影片列表
	Page        int64
	RecordCount int64
}

//新剧推荐模型
type ApiNewInfo struct {
	List []*MovieInfo //影片列表
}

type ApiPageListInfo struct {
	List        []*PageInfo //资讯列表
	Page        int64
	RecordCount int64
}

//资讯详细页模型
type ApiPageDetailInfo struct {
	Info *PageInfo //资讯详情
}

func GetIsEnd(has int64, all int64) string {
	var (
		result string
	)
	if has == all {
		result = "已完结"
	} else {
		result = fmt.Sprintf("%d/%d", has, all)
	}
	return result
}
