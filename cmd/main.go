package main

import (
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/rest"
)

func main() {
	container := injector.Create()

	var (
		ctrl rest.Controller
	)

	err := container.Get(&ctrl)
	if err != nil {
		panic(err)
	}

	err = ctrl.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
