package main

import (
	"myuseek/configs"
	"myuseek/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
