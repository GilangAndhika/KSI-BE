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
	app.Get("/portofolios", controller.GetAllPortofolio)

	app.Post("/orders", controller.CreateOrder)

	app.Get("/profile", controller.GetAllProfile)
	app.Get("/profile/:id", controller.GetProfileByID)
	
	app.Get("/users", controller.GetAllUser)
}
