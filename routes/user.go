package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/woohaen88/database"
	"github.com/woohaen88/models"
)

type User struct {
	// this is not the model User, see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UserSerializer(userModel models.User) User {
	// this function is UserSerializer
	// RDB DBinstance -> struct type
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	responseUser := UserSerializer(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(ctx *fiber.Ctx) error {
	var users []models.User
	database.Database.DB.Find(&users)
	var responseUsers []User

	for _, user := range users {
		responseUser := UserSerializer(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return ctx.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user dose not exist")
	}
	return nil
}

func GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	var user models.User // RDB
	if err != nil {
		return ctx.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	responseUser := UserSerializer(user)
	return ctx.Status(200).JSON(responseUser)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id") // id parsing
	if err != nil {
		return ctx.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User // RDB

	// GET Data
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exists!")
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser // struct

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.DB.Save(&user)

	responseUser := UserSerializer(user) // rdb -> struct
	return ctx.Status(200).JSON(responseUser)

}

func DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON("Please ensure that :id is an integer")
	}

	var user models.User // RDB
	database.Database.DB.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User does not exists!")
	}

	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return ctx.Status(400).JSON("Data가 없음!! ㅅㄱ")
	}

	return ctx.Status(200).JSON("삭제 완료!")

}
