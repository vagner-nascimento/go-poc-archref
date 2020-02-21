package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func main() {
	infra.LogInfo("loading http presentation")
	loadHttpPresentation()
	infra.LogInfo("application running")
	infra.LogInfo("loading subscribers")
	/*
		loadSubscribers MUST be always the last to load because it blocks the app to keep listening to a channel
		that keep consumers connected
	*/
	loadSubscribers()
}
