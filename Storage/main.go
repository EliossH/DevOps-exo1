package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/log", getLog)
	router.POST("/log", postLog)

	router.Run(":8080")
}

func getLog(c *gin.Context) {
	log, err := readlog()
	if err == nil {
		c.String(200, log)
	} else {
		c.String(500, "Error reading log: %v", err)
	}
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

func readlog() (string, error) {
	data, err := os.ReadFile("logs/services.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	return string(data), err
}
