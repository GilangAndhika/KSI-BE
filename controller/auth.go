package controller

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

var secretKey = []byte("Auth") // Gunakan secret key yang aman

// GenerateJWT membuat token JWT
func GenerateJWT(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":    user.ID,  // Menyimpan ID pengguna dalam token
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
		"role":   user.Role,
		"email":  user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
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

	// Generate JWT token
	token, err := GenerateJWT(&input)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating JWT")
	}

	// Set token JWT ke dalam cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // Set cookie expired dalam 24 jam
		HTTPOnly: true, // Hanya bisa diakses oleh server, tidak di JavaScript
		Secure:   false, // Hanya kirim cookie melalui HTTPS (set ke false jika tidak menggunakan HTTPS)
	})

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
		Name:     "jwt",         // Nama cookie
		Value:    t,             // Nilai cookie (JWT)
		Expires:  time.Now().Add(24 * time.Hour), // Masa berlaku cookie
		HTTPOnly: true,          // Akses hanya melalui HTTP (tidak dapat diakses oleh JS)
		Secure:   false,         // Gunakan HTTPS jika true, gunakan false untuk localhost
		SameSite: "Lax",         // Aturan SameSite untuk cookie
	})

	// Return response sukses
	return c.JSON(fiber.Map{
		"message": "Login successful",
	})

}
