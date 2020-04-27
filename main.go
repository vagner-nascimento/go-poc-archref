package main

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/loader"
)

func listenToErrors(errsCh *chan error) {
	fmt.Println("starting to listen app errors")
	for {
		err := <-*errsCh
		fmt.Println("app error", err)
	}
}

// TODO: realise how to close app when it is necessary, when http throws error. for instance
func main() {
	errs := loader.LoadApplication()
	listenToErrors(errs)
}
