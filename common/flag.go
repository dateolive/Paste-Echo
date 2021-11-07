package common

var (
	IpLimitTime = int64(120000) //限制ip的时长 毫秒级别 2分钟

	IpLockCount = 10 //同一个时间ip访问次数

	IpInsideTime = int64(60000) //在多长时间内计算次数  毫秒级别 1分钟

)
