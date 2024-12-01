package controller

import (
	"KSI-BE/repos"

	"github.com/gofiber/fiber/v2"
	"KSI-BE/model"
	// "KSI-BE/config"
)

func GetUserPermissions(c *fiber.Ctx) error {
	role := c.Params("role")
	acl, err := repos.GetPermissionsByRole(role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching permissions")
	}
	return c.JSON(acl)
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	userID, err := repos.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving user")
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": userID,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repos.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func GetAllUser(c *fiber.Ctx) error {
	users, err := repos.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching users")
	}
	return c.JSON(users)
}