package controller

import (
	"context"
	"crud/config"
	"crud/models"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var filter models.User
	err := c.BodyParser(&filter) //takes input from body
	utils.CheckErr(err)
	insert, err := config.Col.InsertOne(context.Background(), filter)
	utils.CheckErr(err)

	return c.Status(200).JSON(insert)
}
