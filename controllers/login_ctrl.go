// Login_ctrl.go
package controllers

import (
	"FD_WebServer/models"
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["ServerHost"] = beego.AppConfig.String("ServerHost")
	this.Data["ServerPort"] = beego.AppConfig.String("ServerPort")
	this.TplName = "login/login.html"
}

func (this *LoginController) Userlogin() {
	fmt.Println("123")
	uname := this.Input().Get("uname")
	fmt.Println(uname)
	pw := this.Input().Get("password")
	fmt.Println(pw)
	userInfo := models.UserLogin(uname, pw)
	fmt.Println(&userInfo)
	if userInfo == nil {

	}

	if pw != userInfo.Pw {

	}

	this.TplName = "tab-subpage-user.html"
}
