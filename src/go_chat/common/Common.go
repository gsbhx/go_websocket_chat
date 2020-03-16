package common

import "github.com/gorilla/websocket"

type UserMessage struct {
	Type        string `json:"type"`
	UserId      int    `json:"user_id"`
	FirstTopic  int    `json:"first_topic"`
	SecondTopic int    `json:"second_topic"`
}

//Redis配置文件结构体
type RedisConf struct {
	Host string //redis host
	Port int    //redis port
}
var (
	UnRegister chan int
	Client     []*websocket.Conn
	RCF=&RedisConf{}
)
