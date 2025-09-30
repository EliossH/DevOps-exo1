package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	createLogFile()
	router := gin.Default()

	router.GET("/log", getLog)
	router.POST("/log", postLog)

	router.Run(":8080")
}

func createLogFile() error {
	file, err := os.OpenFile("services.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()
	return nil
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
	if err != nil {
		c.String(500, "Error reading body: %v", err)
		return
	}

	log := string(body)

	err = writeLog(log)
	if err != nil {
		c.String(500, "Error writing log: %v", err)
		return
	}
}

func readlog() (string, error) {
	data, err := os.ReadFile("services.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	return string(data), err
}

func writeLog(log string) error {
	file, err := os.OpenFile("services.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()

	_, err = file.WriteString(log + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
