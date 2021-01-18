package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title   string `gorm:"column:title;varchat(255);not null" valid:"title"`
	Content string `gorm:"column:content;longtext;not null" valid:"content"`
	UserId  uint64 `gorm:"column:user_id;not null;default:0;index"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
