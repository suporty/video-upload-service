package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const (
	FILE_NAME = "file"
	FILE_DIR  = "./public/"
)

var videoExtMap = map[string]bool{
	".mp4": true,
}

func init() {
	// create default file upload folder
	if _, err := os.Stat(FILE_DIR); os.IsNotExist(err) {
		os.Mkdir(FILE_DIR, 0777)
	}
}

func (s *Server) uploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile(FILE_NAME)
	if err != nil {
		log.Println(err)
		return err
	}

	if !isVideoFile(file.Filename) {
		log.Println(fmt.Sprintf("%s is not a video file", file.Filename))
		return errors.New(fmt.Sprintf("%s is not a video file", file.Filename))
	}

	if err := c.SaveFile(file, fmt.Sprintf("%s%s", FILE_DIR, file.Filename)); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func isVideoFile(filename string) bool {
	ext := filepath.Ext(filename)
	if _, ok := videoExtMap[ext]; !ok {
		return false
	}

	return true
}
