package controller

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/repos"
	// "KSI-BE/model"
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
