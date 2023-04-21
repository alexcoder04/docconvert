package handlers

import "github.com/gin-gonic/gin"

func HandleIndex(c *gin.Context) {
	Templates.ExecuteTemplate(c.Writer, "index.html", gin.H{
		"text": "",
	})
}
