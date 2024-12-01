package controller

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/repos"
	"KSI-BE/model"
)

func CreatePortofolio(c *fiber.Ctx) error {
	var input model.Portofolio

	// Parse body request untuk mendapatkan data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Simpan data portofolio ke database
	portofolioID, err := repos.CreatePortofolio(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving portofolio")
	}

	// Return response sukses dengan portofolio ID
	return c.JSON(fiber.Map{
		"message": "Portofolio created successfully",
		"portofolio_id": portofolioID,
	})
}