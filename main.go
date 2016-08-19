package main

import (
	"github.com/astaxie/beego"
	"github.com/zituocn/M.VMovie/models"
	_ "github.com/zituocn/M.VMovie/routers"
)

func main() {
	beego.AddFuncMap("getisend", models.GetIsEnd)
	beego.Run()
}
