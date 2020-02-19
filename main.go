package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func main() {
	loadHttpPresenter()
	infra.LogInfo("http presentation loaded")
	infra.LogInfo("application running")
	infra.LogInfo("loading subscribers")
	/*
		loadSubscribers MUST be always the last to load because it blocks the app to keep listening to a channel
		that keep consumers connected
	*/
	loadSubscribers()
}
