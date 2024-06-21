package main

import (
	"gorinha/src/db"
	"gorinha/src/routes"
)

func main() {
	db.Init()
	r := routes.SetupRoutes()
	r.Run()
}
