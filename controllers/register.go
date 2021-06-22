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

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) HandleRegister() {
	if (this.Ctx.Input.Method() == "GET") {
		flash := beego.ReadFromRequest(&this.Controller)
		_ = flash
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		this.Data["token"] = token
		this.SetSession("register_token", token)
		this.TplName = "register/register.tpl"
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
				res, err := o.Raw("select Id from myuser where username = ?", user.Username).Exec()
				if err == nil {
				    num, _ := res.RowsAffected()
					if (num == 0) {
						hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
						user.Password = string(hash)
						_, err = o.Insert(&user)
						if(err != nil){
							flash.Error("failed to create user")
							flash.Store(&this.Controller)
							this.Redirect("/register", 302)
						}
						this.Redirect("/login", 302)
					} else {
						flash.Error("username existed!")
						flash.Store(&this.Controller)
						this.Redirect("/register", 302)
					}
				} else {
					flash.Error("failed to create user")
					flash.Store(&this.Controller)
					this.Redirect("/register", 302)
				}
			} else {
				flash.Error("do not hack me")
				flash.Store(&this.Controller)
				this.Redirect("/register", 302)
			}
			
		} else {
			r := ""
			 if valid.HasErrors() {
		        for _, err := range valid.Errors {
		            r += err.Key + ":" + err.Message
		        }
		    }
			flash.Error(r)
			flash.Store(&this.Controller)
			this.Redirect("/register", 302)
		}
	}
}
