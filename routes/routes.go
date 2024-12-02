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
<<<<<<< HEAD
	app.Get("/profile/:id", controller.GetProfileByID)
	
	app.Get("/users", controller.GetAllUser)
	app.Get("/user/:id", controller.GetUserByID)
	app.Post("/user", controller.CreateUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)
=======
>>>>>>> 503a0f8f116eded3b4f850fedfa38474b9328a90
}
