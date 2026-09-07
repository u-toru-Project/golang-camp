package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := http.NewServeMux()

	r.HandleFunc("GET /hello", con.HelloHandler)

	r.HandleFunc("POST /article", con.PostArticleHandler)
	r.HandleFunc("GET /article/list", con.ArticleListHandler)
	r.HandleFunc("GET /article/{id}", con.ArticleDetailHandler)
	r.HandleFunc("POST /article/nice", con.PostNiceHandler)

	r.HandleFunc("POST /comment", con.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
