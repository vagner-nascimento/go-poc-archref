package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

func main() {
	defer data.CloseConnections()
	infra.LogInfo("application running")
	infra.LogInfo("loading subscribers...")
	loadSubscribers() // It MUST be always the last because it keeps listening to a channel to keep consumer connected
}
