package main

import (
	"github.com/MarkTBSS/072_Echo_Server/config"
	"github.com/MarkTBSS/072_Echo_Server/server"
)

func main() {
	conf := config.ConfigGetting()
	server := server.NewEchoServer(conf)
	server.Start()
}
