package main

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/loader"
)

// TODO: realise if really is reading app errors
func listenToErrors(errsCh <-chan error) {
	fmt.Println("starting to listen app errors")
	for {
		err := <-errsCh
		fmt.Println("app error", err)
	}
}

func main() {
	errs := make(chan error)
	go loader.LoadApplication(errs)
	listenToErrors(errs)
}
