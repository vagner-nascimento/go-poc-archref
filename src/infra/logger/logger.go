package logger

import (
	"fmt"
	"time"
)

func getFormattedMessage(msg string) string {
	return fmt.Sprintf("%s - %s", time.Now().Format("02/01/2006 15:04:05"), msg)
}

func Info(msg string) {
	fmt.Println(getFormattedMessage(msg))
}

func Error(msg string, err error) {
	fmt.Println(getFormattedMessage(msg), err)
}
