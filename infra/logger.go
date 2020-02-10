package infra

import "fmt"

func LogInfo(msg string) {
	fmt.Println(msg)
}

func LogError(msg string, err error) {
	fmt.Println(msg+":", err)
}
