package logger

import (
	"fmt"
	"time"
)

func getFormattedMessage(msg string) string {
	return fmt.Sprintf("%s - %s", time.Now().String(), msg)
}

func Info(msg string) {
	fmt.Println(getFormattedMessage(msg))
}

//TODO: realise how to format date
func Error(msg string, err error) {
	fmt.Println(getFormattedMessage(msg), err)
}
