package routers

import (
	"code2/class2ttsx/ttsx/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //登录业务
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin")
    //激活用户
    beego.Router("/active",&controllers.UserController{},"get:ActiveUser")
}
