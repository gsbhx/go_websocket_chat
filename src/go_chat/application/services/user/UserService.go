package user

import (
	"fmt"
	"go_chat/application/services"
	"go_chat/common"
	"go_chat/pool"
	"reflect"
	"strconv"
	"strings"
)

type UserService struct {
	common.CommonFunction
}

func (u *UserService) GetFdByGroup(FirstTopic int, SecondTopic int) []int {
	rk := services.RedisKeyGroupUser
	rk=strings.Replace(rk,"first_topic",strconv.Itoa(FirstTopic),-1)
	rk=strings.Replace(rk,"second_topic",strconv.Itoa(SecondTopic),-1)
	client, _ := new(pool.Pool).GetRedisInstance()
	result,_:=client.Do("GET",rk)
	resultArr:=strings.Split(u.B2S(result),",")
	fmt.Println("resultArr========================================",resultArr)
	fds:=[]int{}
	for _,userid:=range(resultArr){
		if userid!=""{
			fmt.Println("userid=============",userid,reflect.TypeOf(userid))
			rs,_:=client.Do("HGET",services.FdBindUserRedisKey,userid)
			fmt.Println("rs=============",rs,reflect.TypeOf(rs))
			indexInt,_:=strconv.Atoi(u.B2S(rs))
			fmt.Println("indexInt=============",indexInt,reflect.TypeOf(indexInt))

			fds=append(fds,indexInt)
		}
	}
	return fds

}
