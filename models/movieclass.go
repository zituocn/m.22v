package models

//影片类型
//动作/科幻/情景/罪案/惊悚
type MovieClassInfo struct {
	Id          int64
	Name        string `orm:"size(50)"`
	Ename       string `orm:"size(50)"`
	Title       string `orm:"size(200)"`
	Keywords    string `orm:"size(200)"`
	Description string `orm:"size(300)"`
}
