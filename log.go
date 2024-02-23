package helper_GO

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogInfo(logLevel string, message ...any) {
	if logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "debug" {
		m := "Info: " + fmt.Sprintln(message...)
		log.Print(m)
		saveLog("LogInfo.txt", m)
	}
}

func LogWarning(logLevel string, message ...any) {
	if logLevel == "warning" || logLevel == "error" || logLevel == "debug" {
		m := "Warning: " + fmt.Sprintln(message...)
		log.Print(m)
		saveLog("LogWarning.txt", m)
	}
}

func LogError(logLevel string, message ...any) {
	if logLevel == "error" || logLevel == "debug" {
		m := "Error: " + fmt.Sprintln(message...)
		log.Print(m)
		saveLog("LogError.txt", m)
	}
}

func LogDebug(logLevel string, message ...any) {
	if logLevel == "debug" {
		m := "Debug: " + fmt.Sprintln(message...)
		log.Print(m)
		saveLog("LogDebug.txt", m)
	}
}

func saveLog(filename string, message string) {
	//open logfile and append message in a new Line
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}
	defer file.Close()

	if _, err := file.WriteString(time.Now().Format("2006/01/02 15:04:05") + message); err != nil {
		log.Fatal("Failed to write to log file", err)
	}

}
