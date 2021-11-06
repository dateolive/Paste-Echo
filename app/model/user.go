package model

type UserCollection struct {
	Users []User `json:"users"`
}

type User struct {
	UserId int `json:"userid" form:"userid" query:"userid" gorm:"column:userid"`
	UserName  string `json:"username" form:"username" query:"username" gorm:"column:username"`
	Password string `json:"password" form:"password" query:"password" gorm:"column:password"`
	Email string `json:"email" form:"email" query:"email" gorm:"column:email"`
}




