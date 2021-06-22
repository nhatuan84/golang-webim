package controllers

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type ChatController struct {
	beego.Controller
}

func (this *ChatController) HandleSession() {
	chatWith := this.Ctx.Input.Param(":session")
	userName := fmt.Sprint(this.GetSession("logedin"))
	if(userName == ""){
		this.Redirect("/", 302)
	}
	who := strings.Replace(chatWith, userName, "", -1)
	this.Data["who"] = who
	this.Data["session"] = chatWith
	this.TplName = "chat/chat.tpl"
}

func (this *ChatController) StartChat() {
	this.EnableRender = false
	chatWith := this.Ctx.Input.Param(":with")
	userName := fmt.Sprint(this.GetSession("logedin"))
	if(userName == ""){
		this.Redirect("/", 302)
	}
	
	conn, _ := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	who := Who{Username: userName, Conn: conn}
	muGroupChat.Lock()
	groupChat[chatWith] = append(groupChat[chatWith], who)
	muGroupChat.Unlock()

	messages := make(chan string, 4)

	go func() { 
		for{
			msg := <-messages
			if(msg != ""){
				muGroupChat.Lock()
				list := groupChat[chatWith]
				muGroupChat.Unlock()
				for _, user := range list {
			        if err := user.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			        	who := Who{Username: user.Username, Connected: DISCONNECTED}
						notifMessages <- who
		            	return
		        	}
			    }  
			}	
    	}
	}()

    for {
        // Read message from browser
        _, msg, err := conn.ReadMessage()
        if err != nil {
            return
        }
        myMessage := userName + ": " + string(msg)
        messages <- myMessage
    }

}