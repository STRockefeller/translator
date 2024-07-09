package utils

import (
	"fmt"
	"log"
	"os"
)

// Handle Error
func HandleError(err error) {
	if err != nil {
		logToFile(err)
		fmt.Println("An error occurred:", err)
	}
}

// logToFile
func logToFile(err error) {
	f, fileErr := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Println("Failed to open error log file:", fileErr)
		return
	}
	defer f.Close()

	logger := log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(err)
}
