// logger/logger.go

package logger

import (
	"encoding/json"
	"os"
)

func LogToFile(logEntry map[string]interface{}, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logEntryJSON, err := json.Marshal(logEntry)
	if err != nil {
		return err
	}

	formattedLog := string(logEntryJSON) + ","
	_, err = file.Write(append([]byte(formattedLog), '\n'))
	if err != nil {
		return err
	}

	return nil
}
