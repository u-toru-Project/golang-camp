package controllers_test

import (
	"testing"

	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/controllers/services"
	"github.com/yourname/reponame/models"
)

type articleServiceStub struct{}

func (s *articleServiceStub) PostArticleService(article models.Article) (models.Article, error) {
	return article, nil
}

func (s *articleServiceStub) GetArticleListService(page int) ([]models.Article, error) {
	return []models.Article{{ID: 1, Title: "test"}}, nil
}

func (s *articleServiceStub) GetArticleService(articleID int) (models.Article, error) {
	return models.Article{ID: articleID, Title: "test"}, nil
}

func (s *articleServiceStub) PostNiceService(article models.Article) (models.Article, error) {
	return article, nil
}

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	var svc services.ArticleServicer = &articleServiceStub{}
	aCon = controllers.NewArticleController(svc)

	m.Run()
}
