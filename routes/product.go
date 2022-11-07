package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/woohaen88/database"
	"github.com/woohaen88/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

// model -> serializer -> struct
func productSerializer(model models.Product) Product {
	return Product{
		ID:           model.ID,
		Name:         model.Name,
		SerialNumber: model.SerialNumber,
	}
}

func CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&product)

	responseProduct := productSerializer(product)
	return ctx.Status(201).JSON(responseProduct)
}

func GetProducts(ctx *fiber.Ctx) error {
	var products []models.Product

	database.Database.DB.Find(&products)

	var responseProducts []Product
	for _, product := range products {
		responseProduct := productSerializer(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return ctx.Status(200).JSON(responseProducts)
}

func findProduct(model *models.Product, id int) error {
	database.Database.DB.Find(&model, "id = ?", id)
	return nil
}

func GetProduct(ctx *fiber.Ctx) error {
	var product models.Product
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := findProduct(&product, id); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	responseProduct := productSerializer(product)
	return ctx.Status(200).JSON(responseProduct)
}

func PutProduct(ctx *fiber.Ctx) error {
	var product models.Product

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := findProduct(&product, id); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	updateProduct := new(UpdateProduct)

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.DB.Save(&product)

	responseProduct := productSerializer(product)
	return ctx.Status(201).JSON(responseProduct)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	var product models.Product

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := findProduct(&product, id); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	if err := database.Database.DB.Delete(&product).Error; err != nil {
		return ctx.Status(404).JSON(err.Error())
	}

	return ctx.Status(200).JSON("Delete Success!!!!!")
}

