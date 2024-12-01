package controller

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/repos"
	// "KSI-BE/model"
)

func GetAllProfile(c *fiber.Ctx) error {
	// Ambil token dari header
	token := c.Get("Auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Ambil data profile dari database
	profiles, err := repos.GetAllProfile()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching profiles")
	}

	// Return response sukses dengan data profile
	return c.JSON(fiber.Map{
		"profiles": profiles,
	})
}