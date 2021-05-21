package cmd

import (
	"aspect_apps/config"
	"aspect_apps/internal/http/rest"
	"aspect_apps/internal/shared"
	"fmt"
)

func init() {
	fmt.Printf("is run")
}

func StartAPI() {

	var service rest.Service
	domainEvent := shared.NewDomainEventHub()
	conf, err := config.DefaultConfig()

	if err != nil {
		panic(err)
	}
	// Define all service

	// End of defining each service

	// injecting app service Start
	service.Test = "hai"
	// injecting service End

	// Event Driven Service Start

	domainEvent.Start() // Start the Event service
	// Event Driven Service End

	// Run the apps
	rest := rest.NewApp(&service)
	rest.StartApp(fmt.Sprintf(":%s", conf.App.Port))
}
