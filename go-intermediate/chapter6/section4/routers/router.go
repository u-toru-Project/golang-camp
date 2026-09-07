package routers

import (
	"net/http"

	"github.com/yourname/reponame/controllers"
)

func NewRouter(con *controllers.MyAppController) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /hello", con.HelloHandler)

	r.HandleFunc("POST /article", con.PostArticleHandler)
	r.HandleFunc("GET /article/list", con.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", con.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", con.PostNiceHandler)

	r.HandleFunc("POST /comment", con.PostCommentHandler)

	return r
}
