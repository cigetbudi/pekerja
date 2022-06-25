package main

import (
	"pekerja/db"
	"pekerja/routes"
)

func main() {
	db.Init()
	route := routes.Init()

	route.Start(":3435")

}
