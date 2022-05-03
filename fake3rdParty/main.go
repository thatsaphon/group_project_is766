package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/apply/cv", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		email, _ := c.GetPostForm("email")
		fmt.Println(file.Filename)
		fmt.Println(email)
		err := c.SaveUploadedFile(file, "temp/"+file.Filename)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "save file fail",
			})
		}
		c.JSON(200, gin.H{
			"message": "CV is received.",
		})

		time.Now().Format("2006-01-02 150405")
	})

	r.Run(":8081")
}
