package main

import (
	"pekerja/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	route := echo.New()

	route.Start(":6666")

}
