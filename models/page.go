package models

import (
	"time"
)

//页面/专题模型
//手动模式 自定义程度较高的专属页面
type PageInfo struct {
	Id          int64
	Name        string    `orm:"size(200)"`   //中文专题名称
	Ename       string    `orm:"sisze(200)"`  //英语说明  /page/newsyear/
	Content     string    `orm:"size(10000)"` //文章正文 html
	Title       string    `orm:"size(300)"`   //seo标题
	Description string    `orm:"size(500)"`   //seo页面说明
	Keywords    string    `orm:"size(200)"`   //seo关键字
	Status      int64     //状态 -1时，前台不显示 0 为正常
	Views       int64     //浏览量
	Editor      string    `orm:"size(20)"` //责任编辑 显示在页面上
	Addtime     time.Time `orm:"auto_now_add;type(datetime)"`
}
