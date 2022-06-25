package routes

import (
	"pekerja/controllers"
	"pekerja/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	route := echo.New()

	route.GET("/", controllers.FetchAllPersonal)
	route.GET("personal/cari", controllers.FetchAllPersonal)
	route.POST("personal/tambah", controllers.AddPersonal, middleware.IsAuthenticated)
	route.PUT("personal/edit/:email", controllers.EditPersonal, middleware.IsAuthenticated)
	route.DELETE("personal/hapus/:email", controllers.HapusPersonal, middleware.IsAuthenticated)

	route.GET("/genhash/:password", controllers.GenerateHashPass)
	route.POST("/login", controllers.CheckLogin)
	return route
}
