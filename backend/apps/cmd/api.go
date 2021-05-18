package cmd

import (
	"aspect_apps/internal/http/rest"
	"fmt"
)

func init() {
	fmt.Printf("is run")
}

func StartAPI() {

	var service rest.Service

	// Define all service

	// End of defining each service

	// Inject app service
	service.Test = "hai"
	// End Of injecting service

	// Run the apps
	rest := rest.NewApp(&service)
	rest.StartApp(":7000")
}
