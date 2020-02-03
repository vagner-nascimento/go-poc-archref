package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/loader"
)

func main() {
	loader.LoadSubscribers()
	infra.LogInfo("Subscribers loaded")
	infra.LogInfo("Application loaded")
}
