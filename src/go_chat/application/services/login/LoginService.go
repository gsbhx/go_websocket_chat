package services

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/wxnacy/wgo/arrays"
	"go_chat/application/services"
	"go_chat/common"
	"go_chat/pool"
	"reflect"
	"strconv"
	"strings"
)

type returnData struct {
	status int
	msg    string
	data   map[string]interface{}
}

type LoginService struct {
	services.CommonService
	common.CommonFunction
	Index int
}

func (l *LoginService) Register(umsg common.UserMessage) (result returnData) {
	if umsg.FirstTopic == 0 || umsg.SecondTopic == 0 {
		logs.Error("必须要带有一级或二级topic才能通过验证！")
	}
	if umsg.UserId == 0 {
		logs.Error("用户ID不存在，请重试！")
	}
	result = l.saveToRedis(umsg)
	return
}

func (l *LoginService) LogOut() (result *returnData) {
	result = l.removeFromRedis(l.Index)
	return
}
func (l *LoginService) removeFromRedis(index int) (result *returnData) {
	//获取redis实例
	i := strconv.Itoa(index)
	client, _ := new(pool.Pool).GetRedisInstance()
	defer client.Close()
	user_id, _ := client.Do("HGET", services.UserBindRedisKey, i)
	fmt.Println("user_id,,,,,,,", user_id)
	if user_id==nil{
		return
	}
	user_id = l.B2S(user_id)
	fmt.Println("user_id,,,,,,,", user_id, reflect.TypeOf(user_id))
	client.Do("HDEL", services.UserBindRedisKey, i)
	//删除当前用户所对应的分组
	rk := services.RedisKeyUserGroup
	rk = strings.Replace(rk, "fd", i, -1)
	group, _ := client.Do("GET", rk)
	groups := l.B2S(group)
	fmt.Println("group is:", reflect.TypeOf(group), group)
	client.Do("DEL", rk)
	groupArr := strings.Split(groups, "_")
	fmt.Println("groupArr================",reflect.TypeOf(groupArr), groupArr)
	rk = services.RedisKeyGroupUser
	rk = strings.Replace(rk, "first_topic", groupArr[0], -1)
	rk = strings.Replace(rk, "second_topic", groupArr[1], -1)
	userList, _ := client.Do("GET", rk)
	userListArr := strings.Split(l.B2S(userList), ",")
	arrIndex:=arrays.Contains(userListArr,user_id)
	userListArr=append(userListArr[:arrIndex], userListArr[arrIndex+1:]...)
	if len(userListArr)==0{
		client.Do("DEL",rk)
	}else{
		client.Do("SET",rk,strings.Replace(strings.Trim(fmt.Sprint(userListArr), "[]"), " ", ",", -1))
	}
	//删除fd绑定的用户。
	client.Do("HDEL",services.FdBindUserRedisKey,user_id)
	return

}

func (l *LoginService) saveToRedis(umsg common.UserMessage) (result returnData) {
	rk := services.RedisKeyGroupUser
	rk = strings.Replace(rk, "first_topic", strconv.Itoa(umsg.FirstTopic), -1)
	rk = strings.Replace(rk, "second_topic", strconv.Itoa(umsg.SecondTopic), -1)
	fmt.Println("rk==============", rk)
	client, _ := new(pool.Pool).GetRedisInstance()
	defer client.Close()
	userList, err := client.Do("GET", rk)
	if err != nil {
		logs.Error("get user list error:", err)
	}
	userListArr:=strings.Split(l.B2S(userList),",")
	userListArr = append(userListArr, strconv.Itoa(umsg.UserId))
	userListArr = l.UniqueArr(userListArr)
	client.Do("SET", rk, strings.Replace(strings.Trim(fmt.Sprint(userListArr), "[]"), " ", ",", -1))

	//存入用户-组
	rk = services.RedisKeyUserGroup
	rk = strings.Replace(rk, "fd", strconv.Itoa(l.Index), -1)
	fmt.Println("rk===============", rk)
	client.Do("SET", rk, strconv.Itoa(umsg.FirstTopic)+"_"+strconv.Itoa(umsg.SecondTopic))
	//用户对应的fd
	client.Do("HSET", services.UserBindRedisKey, strconv.Itoa(l.Index), strconv.Itoa(umsg.UserId))
	//fd对应的user
	client.Do("HSET", services.FdBindUserRedisKey, strconv.Itoa(umsg.UserId), strconv.Itoa(l.Index))

	return

}
