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

	// Ambil token dari header
	token := c.Get("Auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
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

func GetAllPortofolio(c *fiber.Ctx) error {
	// Ambil token dari header
	token := c.Get("Auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Ambil data portofolio dari database
	portofolios, err := repos.GetAllPortofolio()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching portofolios")
	}

	// Return response sukses dengan data portofolio
	return c.JSON(fiber.Map{
		"portofolios": portofolios,
	})
}