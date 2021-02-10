package main

import (
	"Muse/conf"
	"Muse/controller"
)

func main() {
	conf.Init()
	controller.StartService(controller.Init())
}
