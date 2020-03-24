package main

import "github.com/vagner-nascimento/go-poc-archref/src/loader"

func keepRunning() {
	<-make(chan bool)
}

func init() {
	if err := loader.LoadApplication(); err != nil {
		panic(err)
	}
}

func main() {
	//TODO: Refactor DATA
	keepRunning()
}
