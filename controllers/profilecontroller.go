package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/database"
	"github.com/rinrey24/ptmmlj-be/models"
	"gorm.io/gorm"
)

func GetProfile(c *fiber.Ctx) error {
	var profile []models.Profile
	database.DB.Find(&profile)

	return c.JSON(profile)
}

func ShowProfile(c *fiber.Ctx) error {

	id := c.Params("id")
	var profile models.Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(profile)
}

func CreateProfile(c *fiber.Ctx) error {

	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(profile)
}
