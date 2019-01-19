package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

// https://echo.labstack.com/middleware

// SetupBasic ...
func SetupBasic(e *echo.Echo, apiKey string) {

	// HTTPリクエストのログ出し
	// https://echo.labstack.com/middleware/logger
	// FIXME: GAEなのでそもそもEcho標準ないしフォーマットのみカスタマイズできるミドルウェアは使わない方がよいか？
	e.Use(middleware.Logger())

	// パニックからの復帰
	// https://echo.labstack.com/middleware/recover
	e.Use(middleware.Recover())

	// クロスオリジン対応（ローカルではフロントとサーバとでポート分けて動作確認したりするので）
	// https://echo.labstack.com/middleware/cors
	//if util.IsLocal() {
	e.Use(middleware.CORS())
	//}

	// リクエスト毎にユニークなIDをHTTPヘッダ「X-Request-ID」に積む
	// https://echo.labstack.com/middleware/request-id
	// TODO: どうも値が入らないのでいったん自前で積んでおくことにする
	//e.Use(middleware.RequestID())

	// 以下から保護
	// ・クロスサイトスクリプティング（XSS）攻撃
	// ・コンテンツタイプスニッフィング
	// ・クリックジャック
	// ・安全でない接続およびその他のコードインジェクション攻撃
	// https://echo.labstack.com/middleware/secure
	e.Use(middleware.Secure())

	// TODO: セキュリティ要件としてボディサイズ絞る必要あるかな？
	// https://echo.labstack.com/middleware/body-limit

	// TODO: クロスサイトリクエストフォージェリ対策は必要？
	// https://echo.labstack.com/middleware/csrf

	// TODO: フロントアプリ以外からの呼び出しは想定しないのでAPIキーは使うはず。
	// https://echo.labstack.com/middleware/key-auth
	//if !util.IsLocal() {
	//	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	//		fmt.Printf("get key: %s, apikey: %s\n", key, apiKey)
	//		return key == apiKey, nil
	//	}))
	//}

	// ロールによるアクセス制御は↓で多少なりとも実装が楽にならないか？
	// https://echo.labstack.com/middleware/casbin-auth

	// セッション機能有効化
	// https://echo.labstack.com/middleware/session
	// TODO: 使うかな？
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// TODO: 静的ファイル使うなら。
	// https://echo.labstack.com/middleware/static

}
