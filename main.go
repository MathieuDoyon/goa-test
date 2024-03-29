package main

import (
	"goaTest/app"
	"goaTest/swagger"

	"github.com/raphael/goa"
)

func main() {
	// Create service
	api := goa.New("API")

	// Setup middleware
	api.Use(goa.RequestID())
	api.Use(goa.LogRequest())
	api.Use(goa.Recover())

	// Mount "account" controller
	c := NewAccountController()
	app.MountAccountController(api, c)
	// Mount "bottle" controller
	c2 := NewBottleController()
	app.MountBottleController(api, c2)

	// Mount Swagger spec provider controller
	swagger.MountController(api)

	// Start service, listen on port 8080
	api.ListenAndServe(":8080")
}
