package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gitee.com/wjlkk/media_process/src/handler"
	"gitee.com/wjlkk/media_process/src/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.InitDB()

	r := gin.Default()

	r.POST("/upload", handler.UploadVideo)

	r.Run(":8080")
}
