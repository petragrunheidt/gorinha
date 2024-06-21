package main

import (
	"gorinha/src/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run()
}
