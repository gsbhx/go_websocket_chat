package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	services "go_chat/application/services/login"
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
	Conn  *websocket.Conn
}

func (m *MessageController) GetMessage() {
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
			result := loginService.Register(umsg)
			fmt.Println(result)

			break
		case "message":
			fmt.Println("this is a message")
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
	common.Client = append(common.Client[:m.Index], common.Client[m.Index+1:]...)
	m.Conn.Close()
	fmt.Println("common.client.len:", len(common.Client))
	fmt.Println("common.client:", common.Client)

}
