package requests

import (
	"goblog/app/models/article"

	"github.com/thedevsaddam/govalidator"
)

func ValidateArticleForm(data article.Article) map[string][]string {

	rules := govalidator.MapData{
		"title":   []string{"required", "min:3", "max:40"},
		"content": []string{"required", "min:10"},
	}

	message := govalidator.MapData{
		"title": []string{
			"required:标题不能为空",
			"min:标题长度为3-40",
			"max:标题长度为3-40",
		},
		"content": []string{
			"reuqired:文章内容不能为空",
			"min:文章内容长度必须大于10",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      message,
	}

	return govalidator.New(opts).ValidateStruct()
}
