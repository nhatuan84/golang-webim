package controllers

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Status int
const (
	DISCONNECTED Status = iota + 1
	CONNECTED
)

type Who struct { 
	Username   	string 
	Conn 		*websocket.Conn
	Connected		Status
}

var muGroupChat sync.Mutex
var groupChat = make(map[string][]Who)


var notifMessages = make(chan Who, 4)
var muNotifWho sync.Mutex
var notifWho = make(map[string]Who)

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func processNotif(){
	for {
		notif := <- notifMessages
		muNotifWho.Lock()
		notifWho[notif.Username] = notif
		muNotifWho.Unlock()
		state := " is online"
		if(notif.Connected == CONNECTED){
			state = " is online"
		} else {
			state = " is offline"
		}
		newConnect := "<a href=pair/" + notif.Username + ">" + notif.Username + "</a>" + state
		oldConnect := ""
		if(notif.Connected == CONNECTED){
			oldConnect = notif.Username + " is online"
		} 
		for k,v := range notifWho{
			if (k != "" && k != notif.Username){
				if err := v.Conn.WriteMessage(websocket.TextMessage, []byte(newConnect)); err != nil {
		        	return
		    	}
		    	if(notif.Connected == CONNECTED){
			    	oldConnect += "\n"
					oldConnect += "<a href=pair/" + k + ">" + k + "</a>" + state
				}
			}
		}
		if(notif.Connected == CONNECTED){
			conn := notifWho[notif.Username].Conn
			if err := conn.WriteMessage(websocket.TextMessage, []byte(oldConnect)); err != nil {
	        	return
	    	}
		} else {
			muNotifWho.Lock()
			delete(notifWho, notif.Username);
			muNotifWho.Unlock()
			
			muGroupChat.Lock()
			delete(groupChat, notif.Username);
			muGroupChat.Unlock()

		}
		_ = oldConnect
	}
}

func init() {
	go processNotif()
}