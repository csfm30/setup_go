package router

import (
	"setup_go/api/account"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	apiBackendPrefix := app.Group("/testgo")
	apiRoutes := apiBackendPrefix.Group("/api")
	v1 := apiRoutes.Group("/v1")

	v1.Post("/addAccount", account.AddAccount)
	v1.Get("/getAllAccount", account.GetAccount)
}
