package controller

import (
	"github.com/gofiber/fiber/v2"
	"KSI-BE/repos"
	"KSI-BE/model"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"log"
)

func Login(c *fiber.Ctx) error {
	var input model.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user, err := repos.GetUserByUsername(input.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create token")
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func Register(c *fiber.Ctx) error {
	var input model.User

	// Parse body request untuk mendapatkan data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Validasi email dan nomor telepon
	if input.Username == "" || input.Email == "" || input.Phone == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("All fields are required")
	}

	// Jika role tidak disertakan, set role ke 0 (user)
	if input.Role == 0 {
		input.Role = 0
	} else if input.Role != 1 && input.Role != 2 {
		// Validasi role
		return c.Status(fiber.StatusBadRequest).SendString("Role must be 0, 1, or 2")
	}

	// Hash password sebelum disimpan
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}
	input.Password = string(hashPassword)

	// Simpan data user ke database
	userID, err := repos.CreateUser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
	}

	// Return response sukses dengan user ID
	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": userID,
	})
}