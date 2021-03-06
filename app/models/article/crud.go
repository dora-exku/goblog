package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	db := model.DB.Model(Article{}).Order("id desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	viewData := _pager.Paging()

	var articles []Article
	if err := _pager.Results(&articles); err != nil {
		return articles, viewData, err
	}
	return articles, viewData, nil
}

func GetByUserId(uid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Order("id desc").Where("user_id = ?", uid)
	_pager := pagination.New(r, db, route.Name2URL("users.show"), perPage)
	viewData := _pager.Paging()
	var articles []Article
	if err := _pager.Results(&articles); err != nil {
		return articles, viewData, err
	}
	return articles, viewData, nil
}

func (article *Article) Create() error {
	if err := model.DB.Create(article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}
