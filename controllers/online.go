package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type OnlineController struct {
	beego.Controller
}

func (this *OnlineController) HandlePair() {
	chatWith := this.Ctx.Input.Param(":pair")
	me := fmt.Sprint(this.GetSession("logedin"))
	if(len(me) == 0){
		this.Redirect("/", 302)
	}
	if(chatWith != me){
		pair := chatWith + me

		muNotifWho.Lock()
		conn := notifWho[chatWith].Conn
		if err := conn.WriteMessage(websocket.TextMessage, []byte(me + 
								" want to chat: " + this.Ctx.Request.Host + "/chat/" + pair)); err != nil {
        	return
    	}
		muNotifWho.Unlock()

	    this.Redirect("/chat/" + pair, 302)
	} else {
		this.EnableRender = false
	}
}

func (this *OnlineController) HandleStatus() {
	userName := fmt.Sprint(this.GetSession("logedin"))
	if(len(userName) == 0){
		this.Redirect("/", 302)
	}
	muNotifWho.Lock()
	value, ok := notifWho[userName] 
	muNotifWho.Unlock()
	if(ok){
		value.Conn.Close()
	}
	conn, _ := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	who := Who{Username: userName, Conn: conn, Connected: CONNECTED}
	notifMessages <- who
	this.EnableRender = false
}

func (this *OnlineController) Get() {
	userName := fmt.Sprint(this.GetSession("logedin"))
	if(len(userName) > 0){
		this.Data["who"] = userName
		this.TplName = "online/online.tpl"
	} else {
		this.Redirect("/", 302)
	}
}
