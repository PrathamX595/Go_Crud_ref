package controller

import (
	"context"
	"crud/config"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get(c *fiber.Ctx) error {
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	utils.CheckErr(err)

	filter := bson.M{"_id": objID}
	var user bson.M

	cur, err := config.Col.Find(context.Background(), filter)
	utils.CheckErr(err)

	if cur.Next(context.Background()) {
		err := cur.Decode(&user)
		utils.CheckErr(err)
	}else{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
		
	defer cur.Close(context.Background())
	return c.Status(200).JSON(user)
}

func GetAll(c *fiber.Ctx) error {
	cur, err := config.Col.Find(context.Background(), bson.D{{}})
	utils.CheckErr(err)

	var users []primitive.M
	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		utils.CheckErr(err)
		users = append(users, user)
	}
	defer cur.Close(context.Background())

	return c.Status(200).JSON(users)
}
