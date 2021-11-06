package paste

import (
	"Paste-Echo/config"
)

type LongTimePaste struct {
	*AbstructPaste

}

//成员函数
func (paste *LongTimePaste)Save()error{
	db:=config.NewDb()

	paste.Key = generate(8,&paste)
	paste.Password = hash(paste.Password)
	return db.Table("paste").Create(&paste).Error
}
//成员函数
func (paste *LongTimePaste)Delete()error{
	db:=config.NewDb()
	return db.Table("paste").Delete(&paste).Error
}

//成员函数
func (paste *LongTimePaste)Get(password string)error{
	db:=config.NewDb()
	if err:=db.Table("paste").Where("`key` = ?",paste.GetKey()).Find(&paste).Error;err!=nil{
		return err
	}
	if flag:=paste.checkPassword(password);flag!=nil{
		return flag
	}
	return nil
}