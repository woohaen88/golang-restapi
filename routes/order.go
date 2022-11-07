package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/woohaen88/database"
	"github.com/woohaen88/models"
	"time"
)

type Order struct {
	ID        uint       `json:"id"`
	User      User       `json:"user"`
	Product   Product    `json:"product"`
	CreatedAt *time.Time `json:"order_date"`
}

func OrderSerializer(
	order models.Order,
	user User,
	product Product,
) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func PostOrder(ctx *fiber.Ctx) error {
	var order models.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserID, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(&product, order.ProductID); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&order)

	responseUser := UserSerializer(user)
	responseProduct := productSerializer(product)

	responseOrder := OrderSerializer(order, responseUser, responseProduct)
	return ctx.Status(201).JSON(responseOrder)
}

func GetOrders(ctx *fiber.Ctx) error {
	var orders []models.Order
	database.Database.DB.Find(&orders)

	var responseOrders []Order
	for _, order := range orders {
		var user models.User
		if err := findUser(order.UserID, &user); err != nil {
			return ctx.Status(400).JSON(err.Error())
		}
		responseUser := UserSerializer(user)

		var product models.Product
		if err := findProduct(&product, order.ProductID); err != nil {
			return ctx.Status(400).JSON(err.Error())
		}
		responseProduct := productSerializer(product)

		responseOrder := OrderSerializer(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}

	return ctx.Status(200).JSON(responseOrders)
}

func GetOrder(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var order models.Order
	if err := findOrder(&order, id); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	if err := findUser(order.UserID, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := findProduct(&product, order.ProductID); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	responseUser := UserSerializer(user)
	responseProduct := productSerializer(product)
	responseOrder := OrderSerializer(order, responseUser, responseProduct)
	return ctx.Status(200).JSON(responseOrder)
}

func findOrder(model *models.Order, id int) error {
	database.Database.DB.First(&model, id)
	return nil
}
