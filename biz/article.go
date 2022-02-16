package biz

import (
	"context"
	"errors"
	"github.com/towelong/vgo/dal/model"
	"github.com/towelong/vgo/dal/query"
	"github.com/towelong/vgo/db"
	errors2 "github.com/towelong/vgo/errors"
	"gorm.io/gorm"
)

var Article IArticle = &ArticleService{}

type IArticle interface {
	GetArticleById(ctx context.Context, id uint) (*model.Article, error)
}

type ArticleService struct{}

func (a *ArticleService) GetArticleById(ctx context.Context, id uint) (*model.Article, error) {
	atc := query.Use(db.DB).Article
	article, err := atc.WithContext(ctx).Where(atc.ID.Eq(int32(id))).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors2.NewError(404, "文章未找到")
	}
	return article, nil
}
