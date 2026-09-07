package api

import (
	"database/sql"
	"net/http"

	"github.com/yourname/reponame/api/middlewares"
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"
)

func NewRouter(db *sql.DB) http.Handler {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := http.NewServeMux()

	r.HandleFunc("GET /hello", aCon.HelloHandler)

	r.HandleFunc("POST /article", aCon.PostArticleHandler)
	r.HandleFunc("GET /article/list", aCon.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", aCon.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", aCon.PostNiceHandler)

	r.HandleFunc("POST /comment", cCon.PostCommentHandler)

	return middlewares.LoggingMiddleware(middlewares.AuthMiddleware(r))
}
