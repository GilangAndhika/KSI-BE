package controller

import "fmt"

import (
	"KSI-BE/model"
	"KSI-BE/repos"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

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


func Login(c *fiber.Ctx) error {
	var input model.User

	// Parse body request untuk mendapatkan data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Validasi username dan password
	if input.Username == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username and password are required")
	}

	// Ambil data user berdasarkan username
	user, err := repos.GetUserByUsername(input.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching user")
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Validasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
	}

	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign token
	t, err := token.SignedString([]byte(os.Getenv("LOGIN")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error signing token")
	}

	// Set token in cookies
	c.Cookie(&fiber.Cookie{
		Name:     "Auth",         // Nama cookie
		Value:    t,             // Nilai cookie (JWT)
		Expires:  time.Now().Add(24 * time.Hour), // Masa berlaku cookie
		HTTPOnly: false,          // Akses hanya melalui HTTP (tidak dapat diakses oleh JS)
		Secure:   false,         // Gunakan HTTPS jika true, gunakan false untuk localhost
		SameSite: "Lax",         // Aturan SameSite untuk cookie
	})

	// Set token in header
	c.Set("Auth", t)

	// Set id user in cookies
	c.Cookie(&fiber.Cookie{
		Name:     "ID",         // Nama cookie
		Value:    user.ID,             // Nilai cookie (JWT)
		Expires:  time.Now().Add(24 * time.Hour), // Masa berlaku cookie
		HTTPOnly: false,          // Akses hanya melalui HTTP (tidak dapat diakses oleh JS)
		Secure:   false,         // Gunakan HTTPS jika true, gunakan false untuk localhost
		SameSite: "Lax",         // Aturan SameSite untuk cookie
	})

	// Set role in cookies
	c.Cookie(&fiber.Cookie{
		Name:     "Role",         // Nama cookie
		Value:    fmt.Sprint(user.Role),             // Nilai cookie (JWT)
		Expires:  time.Now().Add(24 * time.Hour), // Masa berlaku cookie
		HTTPOnly: false,          // Akses hanya melalui HTTP (tidak dapat diakses oleh JS)
		Secure:   false,         // Gunakan HTTPS jika true, gunakan false untuk localhost
		SameSite: "Lax",         // Aturan SameSite untuk cookie
	})

	// Return response sukses
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"role":    user.Role,
		"token":   t,
	})

}
