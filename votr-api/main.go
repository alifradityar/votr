package main

import (
	"fmt"
	"net/http"

	"github.com/alifradityar/votr/votr-api/config"
	"github.com/alifradityar/votr/votr-api/handler"
	"github.com/alifradityar/votr/votr-api/router"
	"github.com/facebookgo/inject"
)

func main() {
	conf := config.Get()

	// Setup dependency injection
	var rh handler.Root
	err := inject.Populate(&rh)
	if err != nil {
		fmt.Printf("Error dependency injection: %v", err)
	}

	// Setup router
	r := router.CreateRouter(rh)

	// Serve
	fmt.Printf("Votr started in Port: %s", conf.Port)
	err = http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
