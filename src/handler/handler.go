package handler

import (
	"fmt"
	"net/http"
	"os/exec"

	"gitee.com/wjlkk/media_process/src/model"
	"gitee.com/wjlkk/media_process/src/service"
	"gitee.com/wjlkk/media_process/src/utils"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file to disk
	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate BV number
	bvNumber := utils.GenerateUUID()

	// Use ffprobe to get video information
	ffprobeOutput, err := exec.Command("ffprobe", "-v", "error", "-show_format", "-show_streams", filePath).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse ffprobe output
	videoInfo := service.ParseFFprobeOutput(ffprobeOutput)

	// Save video information to database
	video := model.Video{
		BVNumber: bvNumber,
		FilePath: filePath,
		Info:     videoInfo,
	}

	if err := model.SaveVideo(&video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bv_number": bvNumber,
		"info":      videoInfo,
	})
}
