package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/database"
	"github.com/rinrey24/ptmmlj-be/models"
	"gorm.io/gorm"
)

func GetStakeholder(c *fiber.Ctx) error {
	var data []models.Stakeholder
	database.DB.Find(&data)

	return c.JSON(data)
}

func ShowStakeholder(c *fiber.Ctx) error {

	id := c.Params("id")
	var data models.Stakeholder
	if err := database.DB.First(&data, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(data)
}

func CreateStakeholder(c *fiber.Ctx) error {

	var data models.Stakeholder
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func UpdateStakeholder(c *fiber.Ctx) error {
	var data models.Stakeholder
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.First(data, data.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	// if err := database.DB.Save(&data).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	if err := database.DB.Updates(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Data berhasil diperbarui",
		"data":    data,
	})
}
