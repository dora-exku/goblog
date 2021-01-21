package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
	"goblog/pkg/route"
	"strconv"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)" valid:"password"`

	PasswordComfirm string `gorm:"-" valid:"password_comfirm"`
}

func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}

func (u User) Link() string {
	return route.Name2URL("users.show", "id", strconv.FormatUint(u.ID, 10))
}
