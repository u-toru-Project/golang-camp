# 1.4 標準ライブラリの ServeMux でメソッドを指定しよう

Go 1.22 以降、`http.NewServeMux` は `"GET /hello"` のように HTTP メソッド付きパターンを登録できます。パスパラメータは `{id}` 形式で指定し、`req.PathValue("id")` で取得します。