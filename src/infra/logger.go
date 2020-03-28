package infra

import "fmt"

// TODO: make a better logger
func LogInfo(msgs ...string) {
	fmt.Println(msgs)
}

func LogError(msg string, err error) {
	fmt.Println(msg, err)
}
