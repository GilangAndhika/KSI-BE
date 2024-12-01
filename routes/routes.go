package routes

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/controller"
)

func Setup(app *fiber.App) {
	app.Post("/login", controller.Login)
	app.Get("/permissions/:role", controller.GetUserPermissions)

	app.Post("/register", controller.Register)

	app.Post("/portofolio", controller.CreatePortofolio)

	app.Post("/orders", controller.CreateOrder)
}
