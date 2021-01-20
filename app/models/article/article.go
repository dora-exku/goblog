package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title   string    `gorm:"column:title;varchat(255);not null" valid:"title"`
	Content string    `gorm:"column:content;longtext;not null" valid:"content"`
	UserId  uint64    `gorm:"column:user_id;not null;default:0;index"`
	User    user.User `gorm:"foreignKey:UserId"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func (a Article) CreatedAtDate() string {

	return a.CreatedAt.Format("2006-01-02")
}
