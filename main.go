package main

import (
	"screening/config"
	"screening/routes"
)

func main() {

	config.InitializeDB("./config.env")
	routes.SetupJsonApi()

}
