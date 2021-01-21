package policies

import (
	"goblog/app/models/article"
	"goblog/pkg/auth"
)

func CanModifyArticle(article article.Article) bool {
	return auth.User().ID == article.UserId
}
