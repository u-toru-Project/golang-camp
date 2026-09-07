package main

import (
	"log"
	"net/http"

	"github.com/yourname/reponame/handlers"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("GET /hello", handlers.HelloHandler)

	r.HandleFunc("POST /article", handlers.PostArticleHandler)
	r.HandleFunc("GET /article/list", handlers.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", handlers.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", handlers.PostNiceHandler)

	r.HandleFunc("POST /comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
