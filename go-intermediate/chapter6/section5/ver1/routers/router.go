package routers

import (
	"net/http"

	"github.com/yourname/reponame/controllers"
)

func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /hello", aCon.HelloHandler)

	r.HandleFunc("POST /article", aCon.PostArticleHandler)
	r.HandleFunc("GET /article/list", aCon.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", aCon.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", aCon.PostNiceHandler)

	r.HandleFunc("POST /comment", cCon.PostCommentHandler)

	return r
}
