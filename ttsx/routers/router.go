package routers

import (
	"code2/class2ttsx/ttsx/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
