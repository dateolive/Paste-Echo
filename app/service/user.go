package service

import (
	"Paste-Echo/app/model"
	"Paste-Echo/config"
	"fmt"
)

var(
	UserDao = new(userDao)
)

type userDao struct {

}

func (userDao)GetUserByUserId(userid int)(*model.User,error){
	db:=config.NewDb()

	user:=new(model.User)

	return user,db.Table("user").Where("userid = ?",userid).Find(&user).Error
}

func (userDao)Save(user *model.User)error{
	db:=config.NewDb()
	fmt.Println(*user)
	return db.Table("user").Create(&user).Error
}