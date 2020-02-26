package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func main() {
	infra.LogInfo("loading http presentation")
	loadHttpPresentation()
	infra.LogInfo("application is running")
	infra.LogInfo("loading subscribers")
	loadSubscribers() // <- subscribers MUST be always the last loaded because it blocks the app to keep listening to the amq channels
}
