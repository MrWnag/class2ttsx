package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"github.com/astaxie/beego/orm"
	"code2/class2ttsx/ttsx/models"
)

type UserController struct {
	beego.Controller
}

//展示注册页面
func(this*UserController)ShowRegister(){
	this.TplName = "register.html"
}

//处理注册业务
func (this*UserController)HandleRegister(){
	//获取数据
	userName :=this.GetString("user_name")
	pwd := this.GetString("pwd")
	cpwd :=this.GetString("cpwd")
	email:=this.GetString("email")
	//校验数据
	if userName == "" || pwd == "" || cpwd == "" || email == ""{
		this.Data["errmsg"] = "输入数据不完整，请重新输入！"
		this.TplName = "register.html"
		return
	}

	reg,err :=regexp.Compile(`^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$`)
	if err != nil {
		this.Data["errmsg"] = "正则创建失败！"
		this.TplName = "register.html"
		return
	}
	res := reg.MatchString(email)
	if res == false {
		this.Data["errmsg"] = "邮箱格式不正确，请重新输入！"
		this.TplName = "register.html"
		return
	}
	if pwd != cpwd{
		this.Data["errmsg"] = "两次密码输入不一致，请重新输入！"
		this.TplName = "register.html"
		return
	}

	//处理数据
	//插入操作
	o := orm.NewOrm()
	//获取插入对象
	var user models.User
	//给插入对象赋值
	user.UserName = userName
	user.Pwd = pwd
	user.Email = email
	//执行插入操作
	_,err =o.Insert(&user)
	if err != nil{
		this.Data["errmsg"] = "用户名重复，请重新输入！"
		this.TplName = "register.html"
		return
	}

	//返回数据
	this.Redirect("/login",302)
}

//展示登录页面
func (this*UserController)ShowLogin(){
	this.TplName = "login.html"
}
