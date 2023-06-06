package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jugui93/rest-api/database"
	"github.com/jugui93/rest-api/models"
	"gorm.io/gorm"
)

type HandlersInterface interface {
	ListFacts(c *fiber.Ctx) error
	CreateFact(c *fiber.Ctx) error
	ShowFact(c *fiber.Ctx) error
	UpdateFact(c *fiber.Ctx) error
	DeleteFact(c *fiber.Ctx) error
}

type Handlers struct {}

func NewHandlers() HandlersInterface {
	return &Handlers{}
}

func (h *Handlers) ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func (h *Handlers) CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func (h *Handlers) ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	err := database.DB.Db.Where("id = ?", id).First(&fact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Fact not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fact)
}

func (h *Handlers) UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")
	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Fact not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fact)
}

func (h *Handlers) DeleteFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")
	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Fact not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = database.DB.Db.Where("id = ?", id).Delete(fact).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Fact deleted successfully",
	})
}