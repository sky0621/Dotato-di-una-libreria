package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

// Routing ...
func Routing(e *echo.Echo) {
	InitValidator()

	http.Handle("/", e)

	// 個人の認証を要するWebAPI用のルート
	g := e.Group("/api/v1")

	// 「お知らせ」機能のルーティング("/notices")
	HandleNotice(g)

}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// InitValidator ...
func InitValidator() {
	validate = validator.New()
}

func parse(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	// See document and examples
	// https://github.com/go-playground/validator/blob/v9/README.md#usage-and-documentation
	err := validate.Struct(i)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	for _, err := range err.(validator.ValidationErrors) {
		if err != nil {
			return fmt.Errorf("namespace:%s, field:%s, structNamespace:%s, tag:%s, actualTag:%s, kind:%s, type:%s, value:%s, param:%s",
				err.Namespace(),
				err.Field(),
				err.StructNamespace(),
				err.Tag(),
				err.ActualTag(),
				err.Kind(),
				err.Type(),
				err.Value(),
				err.Param())
		}
	}

	return nil
}
