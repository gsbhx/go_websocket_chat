package common

import "github.com/gorilla/websocket"

type UserMessage struct {
	Type        string `json:"type"`
	UserId      int    `json:"user_id"`
	FirstTopic  int    `json:"first_topic"`
	SecondTopic int    `json:"second_topic"`
	MsgType int `json:"msg_type"`
	MsgContent string `json:"msg_content"`
}

//Redis配置文件结构体
type RedisConf struct {
	Host string //redis host
	Port int    //redis port
}

var (
	UnRegister chan int
	Client     []*websocket.Conn
	RCF        = &RedisConf{}
)

type ReturnData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type MessageContent struct {
	FirstTopic  int    `json:"first_topic"`
	SecondTopic int    `json:"second_topic"`
	MsgType     int    `json:"msg_type"`
	MsgContent  string `json:"msg_content"`
	UserId      int    `json:"user_id"`
}
