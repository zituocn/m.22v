package models

import (
	"time"
)

type DownAddrInfo struct {
	Id      int64     `json:"id"`
	Mid     int64     `json:"mid"`
	Name    string    `orm:"size(20)" json:"name"`
	Hdtvurl string    `orm:"size(500)"json:"hdtvurl"`
	Mkvurl  string    `orm:"size(500)" json:"mkvurl"`
	Ep      int64     `json:"ep"`
	Addtime time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
}
