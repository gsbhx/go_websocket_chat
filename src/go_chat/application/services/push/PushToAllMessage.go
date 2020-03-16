package push

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"go_chat/common"
)

type PushToAllMessage struct {
	Status int
	Fds    []int
	Msg    string
	Data   common.MessageContent
	Index  int
	Result common.ReturnData
}

func (p *PushToAllMessage) update() bool {
	if p.Msg != "" {
		p.Result.Msg = p.Msg
		p.Result.Status=0
	}else{
		p.Result.Data = p.Data
	}
	jsonData, _ := json.Marshal(p.Result)
	if len(p.Fds) == 0 {
		logs.Error("没有需要被推送的用户！")
		return false
	}
	for _, index := range p.Fds {
		conn := common.Client[index]
		conn.WriteMessage(websocket.TextMessage, jsonData)
	}
	return true
}
