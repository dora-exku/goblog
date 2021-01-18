package article

import (
	"goblog/pkg/auth"

	"gorm.io/gorm"
)

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {

	a.UserId = auth.User().ID

	return
}
