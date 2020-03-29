package main

import (
	"github.com/vagner-nascimento/go-poc-archref/src/loader"
)

func init() {
	if err := loader.LoadApplication(); err != nil {
		panic(err)
	}
}

func keepRunning() {
	<-make(chan bool)
}

func main() {
	keepRunning()
}
