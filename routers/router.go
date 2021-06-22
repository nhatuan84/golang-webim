package routers

import (
	"gochat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.RegisterController{}, "get,post:HandleRegister")
    beego.Router("/login", &controllers.LoginController{}, "get,post:HandleRegister")
    beego.Router("/online", &controllers.OnlineController{})
    beego.Router("/pair/:pair([a-z,A-Z,0-9,.]+)", &controllers.OnlineController{}, "get:HandlePair")
    beego.Router("/status", &controllers.OnlineController{}, "get:HandleStatus")
    beego.Router("/chat/:session([a-z,A-Z,0-9,.]+)", &controllers.ChatController{}, "get:HandleSession")
    beego.Router("/startchat/:with([a-z,A-Z,0-9,.]+)", &controllers.ChatController{}, "get:StartChat")

}
