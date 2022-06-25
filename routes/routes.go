package routes

import (
	"pekerja/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	route := echo.New()

	route.GET("personal/cari", controllers.FetchAllPersonal)
	route.POST("personal/tambah", controllers.AddPersonal)
	route.PUT("personal/edit/:email", controllers.EditPersonal)
	route.DELETE("personal/hapus/:email", controllers.HapusPersonal)

	return route
}
