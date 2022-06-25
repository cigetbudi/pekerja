package main

import (
	"os"
	"pekerja/db"
	"pekerja/routes"
)

func main() {
	db.Init()
	route := routes.Init()
	port := os.Getenv("PORT")
	route.Logger.Fatal(route.Start(":" + port))

}
