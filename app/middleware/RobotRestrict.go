package middleware

import (
	"Paste-Echo/common"
	"Paste-Echo/config"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
	"time"
)

//接口防刷
func RobotRestrict(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		ip := GetRealIp(context.Request())
		if ip == "" {
			return common.ErrFrequentRequest.Abort(context)
		}
		fmt.Println("当前ip是" + ip)
		c := config.Poll.Get()
		defer c.Close()
		timeLong, _ := redis.Int64(c.Do("HGet", "ipBlockedMap", ip))

		curTime := time.Now().UnixNano() / 1e6 //毫秒级别

		fmt.Println(curTime)

		if timeLong != 0 {
			//说明没有解封
			if curTime < timeLong {
				return common.ErrFrequentRequest.Abort(context)
			}
			//如果解封后就删除
			c.Do("HDel", "ipBlockedMap", ip)
			c.Do("HDel", "ipCountMap", ip)
			fmt.Println("已经被解封")
		}

		data, _ := redis.Bytes(c.Do("HGet", "ipCountMap", ip))

		tmList := make([]int64, 0)

		json.Unmarshal(data, &tmList)

		fmt.Printf("获取redis中的 del:%v", tmList)
		//如果切片不等于空
		if len(tmList) != 0 {
			for i := 0; i < len(tmList); {
				//不在限制的时间里
				if curTime > tmList[i]+common.IpInsideTime {
					tmList = append(tmList[:i], tmList[i+1:]...)
				} else {
					i++
				}
			}
			fmt.Printf("after del:%v", tmList)
			if len(tmList) <= common.IpLockCount {
				tmList = append(tmList, curTime)
				fmt.Printf("重新赋值的 del:%v", tmList)

				//进行序列化
				v, _ := json.Marshal(tmList)

				c.Do("HDel", "ipCountMap", ip)
				c.Do("HSet", "ipCountMap", ip, v)
			} else {
				//将ip进行限制
				c.Do("HSet", "ipBlockedMap", ip, curTime+common.IpLimitTime)
				fmt.Println("接口已被限制")
				return common.ErrFrequentRequest.Abort(context)
			}

		} else {
			//如果切片为空
			tmList = append(tmList, curTime)
			fmt.Printf("空切片这里 del:%v", tmList)
			//序列化
			v, _ := json.Marshal(tmList)

			c.Do("HSet", "ipCountMap", ip, v)
		}

		return next(context)
	}
}

func GetRealIp(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}
