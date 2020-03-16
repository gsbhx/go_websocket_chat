package conf

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"go_chat/common"
)

func LoadConf() {
	conf, err := config.NewConfig("ini", "config/configure.conf")
	if err != nil {
		fmt.Println("new config failed,err", err)
		return
	}
	common.RCF.Host = conf.String("redis::host")
	common.RCF.Port, err =conf.Int("redis::port")
}
