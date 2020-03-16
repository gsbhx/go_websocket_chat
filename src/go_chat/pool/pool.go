package pool

import (
	"github.com/garyburd/redigo/redis"
	"go_chat/db"
)

type Pool struct {
}

func (p *Pool) GetRedisInstance()(c redis.Conn,err error) {
	rdb := new(db.RedisInstance)
	c,err=rdb.GetInstance()
	return
}
