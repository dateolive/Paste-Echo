package paste

import (
	"Paste-Echo/common"
	"Paste-Echo/config"
	"crypto/md5"
	"fmt"
	"time"
)


type ILongtimePaste interface {
	Save() error
	Get(string) error
	Delete() error
	GetKey() string
	GetContent() string
	GetLang() string
}

//持久化mysql
type AbstructPaste struct {
	Key       string    `json:"key" swaggerignore:"true" gorm:"type:varchar(16);primaryKey"` // 主键:索引
	Lang      string    `json:"lang" example:"plain" gorm:"type:varchar(16)"`                // 语言类型
	Content   string    `json:"content" example:"Hello World!" gorm:"type:mediumtext"`       // 内容，最大长度为 16777215(2^24-1) 个字符
	Password  string    `json:"password" example:"" gorm:"type:varchar(32)"`                 // 密码
	ClientIP  string    `json:"client_ip" swaggerignore:"true" gorm:"type:varchar(64)"`      // 来源 IP
	//UserId    int64     `json:"userid" swaggerignore:"true" gorm:"type:int"`       			// 用户Id
	CreatedAt time.Time `swaggerignore:"true" gorm:"type:timestamp"`                                               // 存储记录的创建时间
}

func (paste *AbstructPaste)GetKey() string{
	return paste.Key
}

func (paste *AbstructPaste)GetContent()string{
	return paste.Content
}

func (paste *AbstructPaste)GetLang() string{
	return paste.Lang
}

func hash(text string)string{
	if text==""{
		return text
	}
	return fmt.Sprintf("%x",md5.Sum([]byte(text)))
}

func (paste *AbstructPaste)checkPassword(password string)*common.ErrorResponse {
	if paste.Password == hash(password){
		return nil
	}
	return common.ErrWrongPassword
}

//判断key是否有冲突
func exist(key string,modle interface{})bool{
	count:=int64(0)
	db:=config.NewDb()
	db.Table("paste").Model(modle).Where("`key` = ?",key).Count(&count)
	return count>0
}
