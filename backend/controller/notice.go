package controller

import (
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/service"
	"net/http"

	"github.com/labstack/echo"
)

// HandleNotice ... "/notices"パスのルーティング
func HandleNotice(g *echo.Group) {
	g.GET("/notices", listNotice)
}

func listNotice(c echo.Context) error {
	ctx := middleware.GetCustomContext(c)
	log := ctx.GetLog().Path("controller/listNotice")
	log.Infow("CONTROLLER__Start")

	responses, err := service.NewNoticeService(ctx).ListNotice()
	if err != nil {
		log.Errorw(err.Error())
		return c.JSON(http.StatusInternalServerError, errorJSON(http.StatusInternalServerError, err.Error()))
	}

	log.Infow("CONTROLLER__End")
	return c.JSON(http.StatusOK, responses)
}
