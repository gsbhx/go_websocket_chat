package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_chat/common"
)

type RedisInstance struct {
}

func (r *RedisInstance) GetInstance () (client redis.Conn,err error) {

	fmt.Printf("%s:%d\n",common.RCF.Host,common.RCF.Port)
	client, err = redis.Dial("tcp", fmt.Sprintf("%s:%d",common.RCF.Host,common.RCF.Port))
	if err != nil {
		fmt.Println("连接Redis出现问题，", err)
		return
	}
	return
}
