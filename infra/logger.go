package infra

import "fmt"

func LogInfo(msgs ...string) {
	fmt.Println(msgs)
}

func LogError(msg string, err error) {
	fmt.Println(msg+":", err)
}
