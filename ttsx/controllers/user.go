package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"github.com/astaxie/beego/orm"
	"code2/class2ttsx/ttsx/models"
	"github.com/astaxie/beego/utils"
	"strconv"
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

	//注册成功的时候发送激活邮件
	config := `{"username":"563364657@qq.com","password":"olyzkdcepleybcag","host":"smtp.qq.com","port":587}`
	emailSend := utils.NewEMail(config)
	emailSend.From = "563364657@qq.com"
	emailSend.To = []string{email}
	emailSend.Subject = "天天生鲜用户激活"
	emailSend.HTML = `<a href="http://192.168.110.81:8080/active?userId=`+strconv.Itoa(user.Id)+`">点击激活</a>`

	emailSend.Send()

	//返回数据
	//this.Redirect("/login",302)
	this.Ctx.WriteString("注册成功，请前往邮箱激活！")
}

//激活用户
func(this*UserController)ActiveUser(){
	//获取用户id
	userId,err := this.GetInt("userId")
	if err != nil{
		this.Data["errmsg"] = "激活失败，请检查网络!"
		this.TplName = "register.html"
		return
	}
	//更新userId对应用户的active字段
	//获取orm对象
	o:= orm.NewOrm()
	//获取更新对象
	var user models.User
	//给更新对象赋值
	user.Id = userId
	//读取一下
	err = o.Read(&user)
	if err != nil{
		this.Data["errmsg"] = "激活失败，用户不存在!"
		this.TplName = "register.html"
		return
	}

	user.Active = 1
	_,err = o.Update(&user)
	if err != nil{
		this.Data["errmsg"] = "激活失败，更新用户出问题了!"
		this.TplName = "register.html"
		return
	}
	this.Redirect("/login",302)
}

//展示登录页面
func (this*UserController)ShowLogin(){

	userName :=this.Ctx.GetCookie("userName")
	if userName != ""{
		this.Data["userName"] = userName
		this.Data["checked"] = "checked"
	}else {
		this.Data["userName"] = ""
		this.Data["checked"] = ""
	}

	this.TplName = "login.html"
}

//处理登录业务
func (this*UserController)HandleLogin(){
	//获取数据
	userName := this.GetString("username")
	pwd := this.GetString("pwd")
	//校验数据
	if userName == "" || pwd == "" {
		beego.Error("输入数据不完整，请重新输入！")
		this.TplName = "login.html"
		return
	}

	//处理数据
	//查询
	o := orm.NewOrm()
	//获取查询对象
	var user models.User
	//给查询条件赋值
	user.UserName = userName
	//查询
	err := o.Read(&user,"UserName")
	if err != nil{
		beego.Error("用户名不存在，请重新输入！")
		this.TplName = "login.html"
		return
	}

	if user.Pwd != pwd{
		beego.Error("用户名密码不匹配，请重新输入！")
		this.TplName = "login.html"
		return
	}

	if user.Active == 0 {
		beego.Error("用户名未激活，请先去目标邮箱激活！")
		this.TplName = "login.html"
		return
	}

	//利用cookie实现记住用户名
	remember :=this.GetString("remember")
	beego.Info("remember = ",remember)
	if remember == "on"{
		this.Ctx.SetCookie("userName",userName,3600)
	}else {
		this.Ctx.SetCookie("userName",userName,-1)
	}

	this.SetSession("userName",userName)

	//返回数据
	this.Redirect("/",302)
}


//退出登录
func(this*UserController)Logout(){
	//删除session
	this.DelSession("userName")
	//跳转页面
	this.Redirect("/",302)
}

//展示用户中心信息页
func(this*UserController)ShowUserCenterInfo(){

	//获取用户名
	userName := this.GetSession("userName")
	this.Data["userName"] = userName.(string)

	this.Layout = "layout.html"
	this.TplName = "user_center_info.html"
}

//展示用户中心订单页
func(this*UserController)ShowUserCenterOrder(){
	this.Layout = "layout.html"
	this.TplName = "user_center_order.html"
}

//展示用户中心地址页
func(this*UserController)ShowUserCenterSite(){
	this.Layout = "layout.html"
	this.TplName = "user_center_site.html"
}