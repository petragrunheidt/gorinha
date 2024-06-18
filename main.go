package main

import (
	"gorinha/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run()
}