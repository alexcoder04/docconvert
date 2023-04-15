package handlers

import "github.com/gin-gonic/gin"

func HandleIndex(c *gin.Context) {
	Template.Execute(c.Writer, gin.H{
		"text": "",
	})
}
