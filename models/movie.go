package models

import (
	"time"
)

//影片模型
type MovieInfo struct {
	Id          int64
	Name        string    `orm:"size(500)"` //影片中文名称
	Ename       string    `orm:"size(100)"` //影片英文名称
	Cid         int64     //分类id
	Cname       string    //分类名称非数据库字段
	Photo       string    `orm:"size(500)"` //海报
	Iphoto      string    `orm:"size(500)"` //推荐到首页的图片
	Actor       string    `orm:"size(500)"` //主演
	Director    string    `orm:"size(500)"` //导演
	Writer      string    `orm:"size(500)"` //编剧
	Language    string    `orm:"size(20)"`  //语言
	Updateweek  int64     //更新星期x
	Playdate    string    `orm:"size(20)"` //开始播放时间
	Content     string    `orm:"size(2000)"`
	Title       string    `orm:"size(500)"`  //seo标题
	Keywords    string    `orm:"size(500)"`  //seo关键字
	Description string    `orm:"size(1000)"` //seo说明
	Views       int64     //浏览量
	Monthviews  int64     //月浏览量
	Status      int64     //影片类型:0为普通 -1为不可见 1为推荐到首页
	Episode     int64     //总集数
	Hasepisode  int64     //已经更新n集
	Isnew       int64     //是否新剧，即新剧推荐
	Addtime     time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	Updatetime  time.Time `orm:"auto_now_add;type(datetime)"` //更新某季下载的时间，用来排序最近更新
	Editor      string    `orm:"size(50)"`                    //责任编辑
	Isend       int64     //是否已更新完结
}
