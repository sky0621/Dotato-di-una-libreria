package controller

import (
	"Dotato-di-una-libreria/backend/controller/form"
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/service"
	"net/http"

	"github.com/labstack/echo"
)

// HandleUser ... "/users"パスのルーティング
func HandleUser(g *echo.Group) {
	g.POST("/users", createUser)
}

func createUser(c echo.Context) error {
	ctx := middleware.GetCustomContext(c)
	log := ctx.GetLog().Path("controller/createUser")
	log.Infow("CONTROLLER__Start")

	user := &form.User{}
	if errUser := parse(c, user); errUser != nil {
		log.Warnw(errUser.Error())
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, errUser.Error()))
	}

	if err := service.NewUser(ctx, c.Request().Context()).CreateUser(user.ParseToDto()); err != nil {
		log.Warnw(err.Error())
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	log.Infow("CONTROLLER__End")
	return c.JSON(http.StatusOK, "{}")
}
