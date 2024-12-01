package controller

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/repos"
	"KSI-BE/model"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
	var input model.Orders

	// Parse body request untuk mendapatkan data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
	}

	// Ambil token dari header
	token := c.Get("Auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Validasi user_id
	if input.UserID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User ID is required")
	}

	// Ambil data user berdasarkan user_id
	user, err := repos.GetUserByID(input.UserID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch user: " + err.Error())
	}

	// Isi data order dengan informasi dari user
	input.GenerateID()              // Generate ID untuk order
	input.FillUserDetails(user)     // Isi username, email, dan phone dari user
	input.OrdersDate = time.Now()   // Set tanggal order dengan waktu saat ini

	// Simpan order ke database
	orderID, err := repos.CreateOrder(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create order: " + err.Error())
	}

	// Return response berhasil
	return c.JSON(fiber.Map{
		"message":  "Order created successfully",
		"order_id": orderID,
	})
}
