package routes

import (
	"pekerja/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	route := echo.New()

	route.GET("personal/cari", controllers.FetchAllPersonal)
	route.POST("personal/tambah", controllers.AddPersonal)

	return route
}
