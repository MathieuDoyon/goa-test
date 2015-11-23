package main

import (
	"goaTest/app"

	"github.com/raphael/goa"
)

// BottleController implements the bottle resource.
type BottleController struct {
	goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service goa.Service) BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	return nil
}

// Delete runs the delete action.
func (c *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	return nil
}

// List runs the list action.
func (c *BottleController) List(ctx *app.ListBottleContext) error {
	res := app.BottleCollection{}
	return ctx.OK(res, "default")
}

// Rate runs the rate action.
func (c *BottleController) Rate(ctx *app.RateBottleContext) error {
	return nil
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	res := &app.Bottle{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *BottleController) Update(ctx *app.UpdateBottleContext) error {
	return nil
}
