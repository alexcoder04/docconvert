package handlers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/gopandoc"
	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	FromType string `form:"fromtype" binding:"required"`
	ToType   string `form:"totype" binding:"required"`
}

func HandleFileUpload(c *gin.Context) {
	var form UploadRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// only allow certain types
	if !friendly.ArrayContains(ALLOWED_FORMATS, form.FromType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'fromtype' parameter"})
		return
	}
	if !friendly.ArrayContains(ALLOWED_FORMATS, form.ToType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'totype' parameter"})
		return
	}

	// get uploaded content
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	defer file.Close()

	if header.Size > 5*1024*1024 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB limit"})
		return
	}

	content := make([]byte, header.Size)
	if _, err := io.ReadFull(file, content); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// convert document
	data, err := gopandoc.BytesToBytes(content, form.FromType, form.ToType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(header.Filename)
	name := header.Filename[:len(header.Filename)-len(extension)]
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+name+"."+form.ToType)
	c.Data(http.StatusOK, "application/octet-stream", data)
}
