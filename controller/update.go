package controller

import (
	"context"
	"crud/config"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(id)
	utils.CheckErr(err)
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"verified": true}}
	res, err := config.Col.UpdateOne(context.Background(), filter, update)
	utils.CheckErr(err)

	return c.Status(200).JSON(res.ModifiedCount)
}
