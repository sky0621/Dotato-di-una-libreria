package middleware

import (
	"Dotato-di-una-libreria/backend/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

// SetupCustom ...
func SetupCustom(e *echo.Echo, appLgr logger.AppLogger, db *gorm.DB) {
	// コンテキストにDB接続情報等を積んで引き回していくためのカスタマイズ
	e.Use(CustomContextMiddleware())

	// 順番がとても大事！
	e.Use(RequestID())
	e.Use(CustomLoggerMiddleware(appLgr))
	e.Use(RelationalDBAccessor(db))

	//e.Use(WithCredentialsMiddleware())
}

// CustomContext ... Cloud SQLアクセッサ等をcontrollerで受け取れるよう、Echoコンテキストを拡張
type CustomContext interface {
	echo.Context
	GetLog() logger.AppLogger
	GetDB() *gorm.DB
}

type customContext struct {
	echo.Context
	log       logger.AppLogger
	db        *gorm.DB
	requestID string
}

// GetLog ...
func (c *customContext) GetLog() logger.AppLogger {
	return c.log
}

// GetDB ...
func (c *customContext) GetDB() *gorm.DB {
	return c.db
}

// GetCustomContext ...
func GetCustomContext(c echo.Context) CustomContext {
	return c.(*customContext)
}

// CustomContextMiddleware ...
func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{
				Context: c,
			}
			return next(cc)
		}
	}
}

// RequestID ...
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			requestID, err := uuid.NewUUID()
			if err == nil {
				cctx.requestID = requestID.String()
			} else {
				cctx.log.Errorw(err.Error())
			}
			return next(cctx)
		}
	}
}

// CustomLoggerMiddleware ...
func CustomLoggerMiddleware(appLgr logger.AppLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			cctx.log = appLgr
			cctx.log.RequestID(cctx.requestID)
			return next(cctx)
		}
	}
}

// RelationalDBAccessor ... Cloud SQLアクセッサを各HTTPリクエスト処理メソッド内で使用可能にするミドルウェア（デコレータ）
func RelationalDBAccessor(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			cctx.db = db
			cctx.db.SetLogger(cctx.log)
			return next(cctx)
		}
	}
}

// WithCredentialsMiddleware ...
func WithCredentialsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			cctx.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			return next(cctx)
		}
	}
}
