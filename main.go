package main

import (
	"Muse/conf"
	"Muse/controller"
)

func main() {
	conf.Init()
	conf.LoadRuleMap()
	controller.StartService(controller.Init())
}
