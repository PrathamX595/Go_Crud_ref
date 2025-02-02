package controller

import (
	"context"
	"crud/config"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, err := primitive.ObjectIDFromHex(id)
	utils.CheckErr(err)
	filter := bson.M{"_id": userId}
	res, err := config.Col.DeleteOne(context.Background(), filter)
	utils.CheckErr(err)

	return c.Status(200).JSON(res)
}

func DeleteAll(c *fiber.Ctx) error {
	res, err := config.Col.DeleteMany(context.Background(), bson.D{{}}, nil)
	utils.CheckErr(err)

	return c.Status(200).JSON(res)
}
