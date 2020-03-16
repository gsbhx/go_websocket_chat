package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	services "go_chat/application/services/login"
	"go_chat/application/services/message"
	"go_chat/application/services/push"
	"go_chat/application/services/user"
	"go_chat/common"
	"time"
)

const (
	pongTime="0001-01-01 00:00:00 +0000 UTC"
	// 读大小限制，这里设置为1024
	maxMessageSize = 1024
)

type MessageController struct {
	Index int
	Conn *websocket.Conn
}

func (m *MessageController) GetMessage() {
	fmt.Println("当前websocket的index为：",m.Index)
	defer m.CloseWebSocket()
	//读大小限制
	m.Conn.SetReadLimit(maxMessageSize)
	//读超时时间 这里设置为永不超时 time.Time{}
	m.Conn.SetReadDeadline(time.Time{})
	// 超时操作
	m.Conn.SetPongHandler(func(string) error { m.Conn.SetReadDeadline(time.Time{}); return nil })
	for {
		_, msg, err := m.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logs.Error("error: %v", err)
			}
			break
		}
		umsg := common.UserMessage{
			Type:        "",
			UserId:      0,
			FirstTopic:  0,
			SecondTopic: 0,
			MsgType: 0,
			MsgContent: "",
		}
		fmt.Println("msg==============", msg)
		err = json.Unmarshal([]byte(msg), &umsg)
		if err != nil {
			logs.Error("json unmarshal error:", err)
			return
		}
		fmt.Println("umsg===============", umsg)
		switch umsg.Type {
		case "login":
			loginService := services.LoginService{
				Index: m.Index,
			}
			messageContent:=common.MessageContent{
				FirstTopic: umsg.FirstTopic,
				SecondTopic: umsg.SecondTopic,
				MsgType: 0,
				MsgContent: fmt.Sprintf("UserId为%v进入了房间",umsg.UserId),
				UserId: umsg.UserId,
			}
			//给自己推送一个注册成功的信息
			PushToSelf:=new(push.PushToAllMessage)
			PushToSelf.Status=0
			PushToSelf.Msg="注册成功！"
			PushToSelf.Fds=[]int{m.Index}
			loginService.Add(PushToSelf)
			//给所有人推送 xxx进来了的信息
			fds:=new(user.UserService).GetFdByGroup(umsg.FirstTopic,umsg.SecondTopic)
			fds=append(fds,m.Index)
			fmt.Println("fds===================",fds)
			PushToAllObj:=new(push.PushToAllMessage)
			PushToAllObj.Status=0;
			PushToAllObj.Data=messageContent
			PushToAllObj.Fds=fds
			loginService.Add(PushToAllObj)

			result := loginService.Register(umsg)
			fmt.Println(result)
			break
		case "message":
			messageContent:=common.MessageContent{
				FirstTopic: umsg.FirstTopic,
				SecondTopic: umsg.SecondTopic,
				MsgType: umsg.MsgType,
				MsgContent: umsg.MsgContent,
				UserId: umsg.UserId,
			}
			fds:=new(user.UserService).GetFdByGroup(umsg.FirstTopic,umsg.SecondTopic)
			fds=append(fds,m.Index)
			fmt.Println("fds===================",fds)
			PushToAllObj:=new(push.PushToAllMessage)
			PushToAllObj.Status=0;
			PushToAllObj.Data=messageContent
			PushToAllObj.Fds=fds
			messageservice:=new(message.MessageService)
			messageservice.Add(PushToAllObj)
			messageservice.Update()


			break
		default:
			fmt.Println("this is default")
			break

		}
	}

}

func (m *MessageController) CloseWebSocket() {
	loginService := services.LoginService{
		Index: m.Index,
	}
	result := loginService.LogOut()
	fmt.Println(result)
	common.Client[m.Index]=nil
	m.Conn.Close()
	fmt.Println("common.client.len:", len(common.Client))
	fmt.Println("common.client:", common.Client)

}
