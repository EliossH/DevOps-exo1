package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/log", getLog)
	router.POST("/log", postLog)

	router.Run("localhost:8080")
}

func getLog(c *gin.Context) {
	c.String(200, "GET log Ok")
}

func postLog(c *gin.Context) {
	body, err := c.GetRawData()
	if err == nil {
		text := string(body)
		c.String(200, "POST log Ok : %s", text)

	} else {
		c.String(500, "Error reading body: %v", err)
	}
}
