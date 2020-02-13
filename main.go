package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func main() {
	// TODO: Implementar interface REST
	infra.LogInfo("application running")
	infra.LogInfo("loading subscribers...")
	loadSubscribers() // It MUST be always the last because it keeps listening to a channel to keep consumer connected
}
