# 8.3 ServeMux によるルーティングが絡んだハンドラユニットテスト

ルーティングテストでは `http.NewServeMux` に `"GET /article/{id}"` を登録し、`ServeHTTP` 経由で `PathValue` が設定されたリクエストをハンドラへ渡します。