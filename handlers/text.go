package handlers

import (
	"net/http"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/gopandoc"
	"github.com/gin-gonic/gin"
)

type TextRequest struct {
	Text   string `form:"text" binding:"required"`
	ToType string `form:"totype" binding:"required"`
}

func HandleTextData(c *gin.Context) {
	var form TextRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// only allow certain types
	if !friendly.ArrayContains(ALLOWED_FORMATS, form.ToType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'totype' parameter"})
		return
	}

	// convert document
	data, err := gopandoc.BytesToBytes([]byte(form.Text), "txt", form.ToType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=yourtext."+form.ToType)
	c.Data(http.StatusOK, "application/octet-stream", data)
}
