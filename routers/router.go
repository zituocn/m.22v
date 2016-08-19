package routers

import (
	"github.com/astaxie/beego"
	"github.com/zituocn/M.VMovie/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/v/:id:int/", &controllers.IndexHandle{}, "*:Detail")
	beego.Router("/m/:cid:int/", &controllers.IndexHandle{}, "*:List")
	beego.Router("/m/:cid:int/:page:int/", &controllers.IndexHandle{}, "*:List")
	beego.Router("/search/:key(.+)/", &controllers.IndexHandle{}, "*:Search")
	beego.Router("/search/:key(.+)/:page:int/", &controllers.IndexHandle{}, "*:Search")
	beego.Router("/today/", &controllers.IndexHandle{}, "*:Today")
	beego.Router("/news/", &controllers.IndexHandle{}, "*:News")
	beego.Router("/news/:page:int/", &controllers.IndexHandle{}, "*:News")
	beego.Router("/article/:id:int/", &controllers.IndexHandle{}, "*:Page")
	beego.Router("/new/", &controllers.IndexHandle{}, "*:New")
	///error handel
	beego.ErrorController(&controllers.HttpErrorHandel{})
}
