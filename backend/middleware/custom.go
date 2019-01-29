package middleware

import (
	"Dotato-di-una-libreria/backend/logger"
	"net/http"

	firebase "firebase.google.com/go"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

// SetupCustom ...
func SetupCustom(e *echo.Echo, appLgr logger.AppLogger, db *gorm.DB, firebaseApp *firebase.App) {
	// コンテキストにDB接続情報等を積んで引き回していくためのカスタマイズ
	e.Use(customContextMiddleware())

	// 順番がとても大事！
	e.Use(requestIDMiddleware())
	e.Use(customLoggerMiddleware(appLgr))
	e.Use(gormDBMiddleware(db))
	e.Use(firebaseAppMiddleware(firebaseApp))
}

// CustomContext ... Cloud SQLアクセッサ等をcontrollerで受け取れるよう、Echoコンテキストを拡張
type CustomContext interface {
	echo.Context
	GetLog() logger.AppLogger
	GetDB() *gorm.DB
	GetFirebaseApp() *firebase.App
	GetRequest() *http.Request
}

type customContext struct {
	echo.Context
	log         logger.AppLogger
	db          *gorm.DB
	firebaseApp *firebase.App
	requestID   string
}

// GetLog ...
func (c *customContext) GetLog() logger.AppLogger {
	return c.log
}

// GetDB ...
func (c *customContext) GetDB() *gorm.DB {
	return c.db
}

// GetFirebaseApp ...
func (c *customContext) GetFirebaseApp() *firebase.App {
	return c.firebaseApp
}

// GetRequest ...
func (c *customContext) GetRequest() *http.Request {
	return c.GetRequest()
}

// GetCustomContext ...
func GetCustomContext(c echo.Context) CustomContext {
	return c.(*customContext)
}

func customContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{
				Context: c,
			}
			return next(cc)
		}
	}
}

func requestIDMiddleware() echo.MiddlewareFunc {
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

func customLoggerMiddleware(appLgr logger.AppLogger) echo.MiddlewareFunc {
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

func gormDBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
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

func firebaseAppMiddleware(firebaseApp *firebase.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cctx, ok := c.(*customContext)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
			cctx.firebaseApp = firebaseApp
			return next(cctx)
		}
	}
}
