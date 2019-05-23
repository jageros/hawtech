package routers

import (
	"hawtech/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/case", &controllers.CaseController{})
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/newsDetail", &controllers.NewsDeyailController{})
	beego.Router("/product", &controllers.ProductController{})
	beego.Router("editor", &controllers.EditorController{})
}
