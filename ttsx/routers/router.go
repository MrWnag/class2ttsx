package routers

import (
	"code2/class2ttsx/ttsx/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/goods/*",beego.BeforeExec,filterFunc)
    beego.Router("/", &controllers.GoodsController{},"get:ShowIndex")
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //登录业务
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    //激活用户
    beego.Router("/active",&controllers.UserController{},"get:ActiveUser")
    //退出登录
    beego.Router("/logout",&controllers.UserController{},"get:Logout")
    //用户中心信息页
    beego.Router("/goods/userCenterInfo",&controllers.UserController{},"get:ShowUserCenterInfo")
    //用户中心订单页
    beego.Router("/goods/userCenterOrder",&controllers.UserController{},"get:ShowUserCenterOrder")
    //用户中心地址页
    beego.Router("/goods/userCenterSite",&controllers.UserController{},"get:ShowUserCenterSite")
}

func filterFunc(ctx*context.Context){
	//获取session
	userName :=ctx.Input.Session("userName")
	if userName == nil{
		ctx.Redirect(302,"/login")
		return
	}
}