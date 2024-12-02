package controller

import (
	"KSI-BE/model"
	"KSI-BE/repos"
	"os"
    "path/filepath"

	// "KSI-BE/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)


func CreatePortofolio(c *fiber.Ctx) error {
    // Cek apakah folder uploads ada, jika tidak, buat folder tersebut
    uploadsDir := "./uploads"
    if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
        err := os.MkdirAll(uploadsDir, os.ModePerm)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Failed to create uploads directory: " + err.Error())
        }
    }

    // Parse multipart form data
    form, err := c.MultipartForm()
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid form data: " + err.Error())
    }

    // Ambil data dari form
    userID := form.Value["id_user"][0]
    if userID == "" {
        return c.Status(fiber.StatusBadRequest).SendString("User ID is required")
    }

    // Ambil data user berdasarkan user_id
    user, err := repos.GetUserByID(userID)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).SendString("User not found")
        }
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch user: " + err.Error())
    }

    // Ambil data desain dari form
    designImage := form.File["design_image"][0]
    designTitle := form.Value["design_title"][0]
    designDescription := form.Value["design_description"][0]
    designType := form.Value["design_type"][0]

    // Cek apakah file desain image ada
    if designImage == nil || designTitle == "" || designDescription == "" || designType == "" {
        return c.Status(fiber.StatusBadRequest).SendString("All design fields are required")
    }

    // Simpan file desain image ke folder uploads
    err = c.SaveFile(designImage, filepath.Join(uploadsDir, designImage.Filename))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to save design image: " + err.Error())
    }

    // Isi data portofolio dengan informasi dari user dan data desain
    var input model.Portofolio
    input.GenerateID()
    input.FillUserDetails(user)
    input.DesignImage = designImage.Filename  // Menyimpan nama file gambar desain
    input.DesignTitle = designTitle
    input.DesignDescription = designDescription
    input.DesignType = designType

    // Simpan portofolio ke database
    portofolioID, err := repos.CreatePortofolio(&input)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to create portofolio: " + err.Error())
    }

    // Return response berhasil
    return c.JSON(fiber.Map{
        "message":       "Portofolio created successfully",
        "portofolio_id": portofolioID,
    })
}




func GetAllPortofolio(c *fiber.Ctx) error {
	// Ambil data portofolio dari database
	portofolios, err := repos.GetAllPortofolio()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching portfolios")
	}

	// Return success response with all portfolios
	return c.JSON(fiber.Map{
		"portofolios": portofolios,
	})
}
