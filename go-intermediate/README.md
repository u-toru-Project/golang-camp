# APIを作りながら進むGo中級者への道
技術書典13にて頒布した同人誌[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx)のサンプルコードです。

無料試し読みは[こちら](https://hsaki.booth.pm/items/4142313)から

## Go 1.26 向けの変更点
原著サンプルは Go 1.17 と gorilla/mux を前提としていました。このディレクトリでは次のように現行 Go へ合わせています。

- 言語バージョンを Go 1.26 に更新
- ルーティングを Go 1.22 以降の標準 `net/http.ServeMux` に置き換え（`GET /article/{id}` と `Request.PathValue`）
- gorilla/mux への依存を削除
- MySQL 公式イメージを 5.7（EOL）から 8.4 へ更新
- `go-sql-driver/mysql` と `google.golang.org/api` を現行版へ更新

第6章 6.1 では、サービス層のメソッド化に合わせてハンドラも `Handler` 構造体経由で依存注入する形に更新しています（6.2 のコントローラ導入への橋渡し）。

## もくじ
- [第1章 HTTPサーバー](./chapter1/README.md)
- [第2章 構造体とjsonの扱い方](./chapter2/README.md)
- [第3章 データベースの扱い方](./chapter3/README.md)
- [第4章 ユニットテスト (基礎編)](./chapter4/README.md)
- [第5章 サービス層の作成](./chapter5/README.md)
- [第6章 アーキテクチャ大改装](./chapter6/README.md)
- [第7章 エラー処理](./chapter7/README.md)
- [第8章 ユニットテスト (応用編)](./chapter8/README.md)
- [第9章 ミドルウェアによるロギング](./chapter9/README.md)
- [第10章 並行処理](./chapter10/README.md)
- [第11章 contextパッケージの導入](./chapter11/README.md)
- [第12章 ユーザー認証](./chapter12/README.md)
