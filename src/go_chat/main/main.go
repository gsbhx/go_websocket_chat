package main

import (
	"go_chat/conf"
	"go_chat/ws"
	"net/http"
)



func main() {
	//读取配置文件
	conf.LoadConf()
	http.HandleFunc("/ws", ws.IndexHandler)
	http.ListenAndServe("0.0.0.0:9999", nil)
}
