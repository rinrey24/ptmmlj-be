package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/database"
	"github.com/rinrey24/ptmmlj-be/models"
	"gorm.io/gorm"
)

func GetArticle(c *fiber.Ctx) error {
	var data []models.Article
	database.DB.Find(&data)

	//return c.JSON(data)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    data,
	})
}

func ShowArticle(c *fiber.Ctx) error {

	id := c.Params("id")
	var data models.Article
	if err := database.DB.First(&data, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    data,
	})
}

func GetActiveArticle(c *fiber.Ctx) error {
	var data []models.Article
	database.DB.Where("is_active = ?", true).Find(&data)

	//return c.JSON(data)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data ditemukan",
		"data":    data,
	})
}

func CreateArticle(c *fiber.Ctx) error {

	var data models.Article
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data berhasil dibuat",
		"data":    data,
	})
}

func UpdateArticle(c *fiber.Ctx) error {
	var data models.Article
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := database.DB.Updates(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Data berhasil diperbarui",
		"data":    data,
	})
}
