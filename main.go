package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
	"github.com/vagner-nascimento/go-poc-archref/infra/docker"
)

func init() {
	docker.WaitForInfra()
}

func main() {
	defer data.CloseConnections()
	// TODO: Implementar interface REST
	infra.LogInfo("application running")
	infra.LogInfo("loading subscribers...")
	loadSubscribers() // It MUST be always the last because it keeps listening to a channel to keep consumer connected
}
