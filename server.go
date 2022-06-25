package main

import (
	"os"
	"pekerja/db"
	"pekerja/routes"
)

func main() {
	db.Init()
	route := routes.Init()

	// route.Start(":3435")
	port := os.Getenv("PORT")
	route.Logger.Fatal(route.Start(":" + port))

}
