package controllers

import (
	"net/http"
	"time"

	"DS.HomeExchange.Homes/configs"
	"DS.HomeExchange.Homes/models"
	"DS.HomeExchange.Homes/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var homeCollection *mongo.Collection = configs.GetCollection(configs.DB, "homes")
var validate = validator.New()

func CreateHome(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var home models.CreateHomeRequest
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&home); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// use the validator library to validate the required fields
	if validationError := validate.Struct(&home); validationError != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationError.Error()}})
	}

	newHome := models.Home{
		Id:                       primitive.NewObjectID(),
		OwnerId:                  home.OwnerId,
		Address:                  home.Address,
		Description:              home.Description,
		NumberOfRequestsRejected: 0,
	}

	result, err := homeCollection.InsertOne(ctx, newHome)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseDto{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Query the database for the newly inserted document
	var createdHome models.Home
	err = homeCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&createdHome)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseDto{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	array := []models.Home{createdHome}
	return c.Status(http.StatusCreated).JSON(responses.ResponseDto{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": array}})
}

func GetAllHomes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var homes []models.Home
	cursor, err := homeCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseDto{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var home models.Home
		cursor.Decode(&home)
		homes = append(homes, home)
	}

	if homes == nil {
		homes = []models.Home{}
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": homes}})
}

// Get a single home by id given in the query param
func GetHome(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	homeId, error := primitive.ObjectIDFromHex(c.Params("id"))
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": error.Error()}})
	}

	var home models.Home
	err := homeCollection.FindOne(ctx, bson.M{"home_id": homeId}).Decode(&home)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	array := []models.Home{home}
	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": array}})
}

// Get all homes owned by a user given owner_id in the query param
func GetHomesByOwner(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ownerId := c.Params("id")

	var homes []models.Home
	cursor, err := homeCollection.Find(ctx, bson.M{"ownerId": ownerId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseDto{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var home models.Home
		cursor.Decode(&home)
		homes = append(homes, home)
	}

	if homes == nil {
		homes = []models.Home{}
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": homes}})
}

// Update a home by id given together with the updated data in the home object given in request body
func UpdateHome(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var home models.Home
	if err := c.BodyParser(&home); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	_, err := homeCollection.UpdateOne(ctx, bson.M{"home_id": home.Id}, bson.M{"$set": home})
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	array := []models.Home{home}
	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": array}})
}

// Delete a home by id given in the query param
func DeleteHome(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	homeId, error := primitive.ObjectIDFromHex(c.Params("id"))
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": error.Error()}})
	}

	_, err := homeCollection.DeleteOne(ctx, bson.M{"home_id": homeId})
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": nil}})
}

// Increase the number of requests rejected for a home by id given in the query param and delete the home if the number of requests rejected is greater than 3
func RejectHome(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	homeId, error := primitive.ObjectIDFromHex(c.Params("id"))
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": error.Error()}})
	}

	var home models.Home
	var err = homeCollection.FindOne(ctx, bson.M{"home_id": homeId}).Decode(&home)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	home.NumberOfRequestsRejected += 1
	if home.NumberOfRequestsRejected > 3 {
		_, err = homeCollection.DeleteOne(ctx, bson.M{"home_id": homeId})
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	} else {
		_, err = homeCollection.UpdateOne(ctx, bson.M{"home_id": homeId}, bson.M{"$set": home})
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	array := []models.Home{home}
	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": array}})
}

// Perform a home exchange between two users
func PerformHomeExchange(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var request models.SuccesfulHomeExchangeRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Get the homes of the two users
	var fromUserHome models.Home
	var toUserHome models.Home

	fromUserHomeId, err := primitive.ObjectIDFromHex(request.FromUserHomeId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	toUserHomeId, err := primitive.ObjectIDFromHex(request.ToUserHomeId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseDto{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	err = homeCollection.FindOne(ctx, bson.M{"home_id": fromUserHomeId}).Decode(&fromUserHome)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	err = homeCollection.FindOne(ctx, bson.M{"home_id": toUserHomeId}).Decode(&toUserHome)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Perform the home exchange
	fromUserHome.OwnerId = request.ToUserId
	toUserHome.OwnerId = request.FromUserId

	_, err = homeCollection.UpdateOne(ctx, bson.M{"home_id": fromUserHomeId}, bson.M{"$set": fromUserHome})
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	_, err = homeCollection.UpdateOne(ctx, bson.M{"home_id": toUserHomeId}, bson.M{"$set": toUserHome})
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ResponseDto{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseDto{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": nil}})
}

