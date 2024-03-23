package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/database"
	"github.com/rinrey24/ptmmlj-be/models"
	"gorm.io/gorm"
)

func GetProfile(c *fiber.Ctx) error {
	var profile []models.Profile
	database.DB.Find(&profile)

	//return c.JSON(profile)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    profile,
	})
}

func ShowProfile(c *fiber.Ctx) error {

	id := c.Params("id")
	var profile models.Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    profile,
	})
}

func ShowActiveProfile(c *fiber.Ctx) error {

	id := c.Params("id")
	var profile models.Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    profile,
	})
}

func CreateProfile(c *fiber.Ctx) error {

	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data berhasil dibuat",
		"data":    profile,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := database.DB.Updates(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data berhasil diperbarui",
		"data":    profile,
	})
}
