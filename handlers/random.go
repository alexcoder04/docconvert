package handlers

import "github.com/gin-gonic/gin"

func HandleRandom(c *gin.Context) {
	Templates.ExecuteTemplate(c.Writer, "random.html", gin.H{
		"text": "",
	})
}
