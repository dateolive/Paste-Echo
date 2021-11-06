package config

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var Poll *redis.Pool


func InitRedisPoll() error{

	Poll= &redis.Pool{
		MaxIdle: 16, //最初的连接数量
		MaxActive: 0,//最大连接数量 不确定用0 按需分配
		IdleTimeout: 300,//连接关闭300秒
		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",viper.GetString("redis.addr"))
			return c,err
		},
	}

	return nil
}

func CloseRedisPoll(){
	if Poll!=nil{
		Poll.Close()
	}
}
