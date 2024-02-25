package helper_GO

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogInfo(logLevel string, message ...any) bool {
	if logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "debug" {
		m := LogPrint("Info", message...)
		SaveLog("LogInfo.txt", m)
		return true
	}
	return false
}

func LogWarning(logLevel string, message ...any) bool {
	if logLevel == "warning" || logLevel == "error" || logLevel == "debug" {
		m := LogPrint("Warning", message...)
		SaveLog("LogWarning.txt", m)
		return true
	}
	return false
}

func LogError(logLevel string, message ...any) bool {
	if logLevel == "error" || logLevel == "debug" {
		m := LogPrint("Error", message...)
		SaveLog("LogError.txt", m)
		return true
	}
	return false
}

func LogDebug(logLevel string, message ...any) bool {
	if logLevel == "debug" {
		m := LogPrint("Debug", message...)
		SaveLog("LogDebug.txt", m)
		return true
	}
	return false
}

func LogPrint(name string, message ...any) string {
	m := name + ": " + fmt.Sprintln(message...)
	log.Print(m)
	return m
}

func SaveLog(filename string, message string) {
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
