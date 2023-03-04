package main

import (
	"github.com/suporty/video-upload-service/config"
	"github.com/suporty/video-upload-service/server"
)

func main() {
	s := server.NewServer(config.GetPort())
	s.Run()
}
