package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"crypto/md5"
	"time"
	"io"
	"strconv"
	"text/template"
	"gochat/models"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) HandleRegister() {
	if (this.Ctx.Input.Method() == "GET") {
		this.SetSession("logedin", "")
		flash := beego.ReadFromRequest(&this.Controller)
		_ = flash
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		this.Data["token"] = token
		this.SetSession("register_token", token)
		this.TplName = "login/login.tpl"
	} else {
		flash := beego.NewFlash()
		user := models.User{}
		if err := this.ParseForm(&user); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
		}
		storedToken := this.GetSession("register_token")
		token := template.HTMLEscapeString(this.GetString("token"))
		valid := validation.Validation{}
		isValid, _ := valid.Valid(user)
		if(isValid){
			if(storedToken == token){
				o := orm.NewOrm()
				dbUser := models.User{Id: -1, Username: "", Password: ""}
				err := o.Raw("select * from myuser where username = ?", user.Username).QueryRow(&dbUser)
				if err == nil {
					if (dbUser.Id == -1) {
						flash.Error("not matched!")
						flash.Store(&this.Controller)
						this.Redirect("/login", 302)
					} else {
						err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
						if(err == nil){
							this.SetSession("logedin", user.Username)
							this.Redirect("/online", 302)
						}
						flash.Error("wrong password!")
						flash.Store(&this.Controller)
						this.Redirect("/login", 302)
					}
				}
			} else {
				flash.Error("do not hack me")
				flash.Store(&this.Controller)
				this.Redirect("/login", 302)
			}
			
		} else {
			flash.Error("not matched!")
			flash.Store(&this.Controller)
			this.Redirect("/login", 302)
		}
	}
}
