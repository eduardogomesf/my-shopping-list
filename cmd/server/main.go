package main

import (
	config "github.com/eduardogomesf/shopping/configs"
	webserver "github.com/eduardogomesf/shopping/internal/infra/web"
)

func main() {
	conf := config.LoadConfig(".")
	ws := webserver.NewWebServer(conf.APPPort)
	ws.Start()
}
