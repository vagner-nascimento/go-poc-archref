package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/presenter"
)

func loadSubscribers() {
	repository.SubscribeAllConsumers()
}

func loadHttpPresenter() {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	router := presenter.HttpRoutes()
	if err := chi.Walk(router, walkFunc); err != nil {
		infra.LogError("cannot load http presentation", err)
		return
	}

	go http.ListenAndServe(":3000", router)
}
