package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexcoder04/friendly/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Img2TextRequest struct {
	Language string `form:"language" binding:"required"`
}

func HandleImg2Text(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check uploaded file
	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is not an image"})
		return
	}
	if !friendly.ArrayContains(ALLOWED_IMAGE_FORMATS, filepath.Ext(file.Filename)[1:]) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file format"})
		return
	}

	// parse request
	var form Img2TextRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save file
	id := uuid.NewString()
	filepath := fmt.Sprintf("/tmp/%s.%s", id, filepath.Ext(file.Filename))
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// run detection
	err = friendly.Run([]string{"tesseract", "-l", form.Language, filepath, "/tmp/" + id}, "")
	if err != nil {
		os.Remove(filepath)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := os.ReadFile("/tmp/" + id + ".txt")
	if err != nil {
		os.Remove(filepath)
		os.Remove("/tmp/" + id + ".txt")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Template.Execute(c.Writer, gin.H{
		"text": string(data),
	})
}
