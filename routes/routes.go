package routes

import (
	"KSI-BE/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/login", controller.Login)
	app.Get("/permissions/:role", controller.GetUserPermissions)

	app.Post("/register", controller.Register)

	app.Post("/portofolio", controller.CreatePortofolio)

	app.Post("/orders", controller.CreateOrder)

	app.Get("/profile", controller.GetAllProfile)
	app.Get("/users", controller.GetAllUser)
}
