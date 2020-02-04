package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/loader"
)

func main() {
	infra.LogInfo("Application running")
	infra.LogInfo("Loading subscribers...")
	loader.LoadSubscribers()
}
