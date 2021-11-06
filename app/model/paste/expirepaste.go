package paste

import (
	"Paste-Echo/config"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

//redis存储
type ExpireTimePaste struct{
	*AbstructPaste
	ExpireSecond   uint64 `json:"expire_second"` // 过期时间
}

//成员函数
func (paste *ExpireTimePaste)Save()error{
	c:=config.Poll.Get()
	defer c.Close()
	paste.Key = generate(8,&paste)
	paste.Password = hash(paste.Password)
	str,_:=json.Marshal(paste)

	if _,err:=c.Do("Set",paste.GetKey(),string(str),"EX",paste.ExpireSecond);err!=nil{
		return err
	}
	return nil
}

//成员函数
func (paste *ExpireTimePaste)Get(password string)error{
	c:=config.Poll.Get()
	defer c.Close()

	body,err:=redis.String(c.Do("Get",paste.GetKey()))
	if err!=nil{
		return err
	}

	json.Unmarshal([]byte(body),&paste)

	if flag:=paste.checkPassword(password);flag!=nil{
		return flag
	}
	return nil
}