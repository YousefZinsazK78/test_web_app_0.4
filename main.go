package main

import (
	"github.com/yousefzinsazk78/test_web_app_0.4/server"
	"github.com/yousefzinsazk78/test_web_app_0.4/server/config"
)

func main() {

	localServer := server.New()
	localServer.Serve(config.ServerPort)
}
