package models

type User struct {
	Id     		int     `form:"-"`
	Username   	string  `form:"username" valid:"MinSize(5);MaxSize(20)"`
	Password 	string	`form:"password" valid:"MinSize(5)"`
	Email 		string 	`form:"email"`
}

func (user *User) TableName() string {
	return "myuser"
}