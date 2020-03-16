package services

type ReturnData struct {
	status int
	msg    string
	data   map[string]interface{}
}

var (
	RedisKeyGroupUser  = "ws_topic_first_topic_second_topic" //分组下对应的用户
	RedisKeyUserGroup  = "ws_user_fd"                        //用户对应的分组
	UserBindRedisKey   = "ws_user_bind_fd_redis_key"         //fd绑定用户的redis_key fd=》user
	FdBindUserRedisKey = "ws_fd_bind_user_redis_key"         //用户绑定fd的redis_key user=>fd
	RedisExpireTime    = 86400
)

type CommonService struct {
	result                 *ReturnData

}