package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wxnacy/wgo/arrays"
	"go_chat/application/controllers"
	"go_chat/common"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type Clients struct {
	conn map[*websocket.Conn]bool

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	fmt.Println(&conn)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("handshake success %v\n",&conn)
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("handshake success: %v", &conn)));
	common.Client =append(common.Client,conn)
	index:=arrays.Contains(common.Client,conn)
	fmt.Println("index is :",index)
	ctls:=controllers.MessageController{
		index,conn,
	}
	go ctls.GetMessage()
}

