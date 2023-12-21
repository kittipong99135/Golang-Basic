package controllers

import (
	"golang_workshop/database"
	"golang_workshop/models"

	"github.com/gofiber/fiber/v2"
)

func CrateUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).JSON(err.Error())
	}
	db.Create(&user)

	return c.JSON(user)
}

func ReadUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user []models.Users

	db.Find(&user)

	return c.Status(200).JSON(user)
}

func ReadUsers(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.Users

	db.Find(&user, "id = ?", id)

	return c.Status(200).JSON(user)
}

func SearchUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := c.Query("search")
	var user models.Users

	result := db.Find(&user, "employee_id = ? OR name = ? OR lastname = ?", search, search, search)
	if result.RowsAffected == 0 {
		return c.Status(400).SendString("search  " + string(search))
	}

	return c.Status(200).JSON(user)
}

func GenerationUser(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.Users
	db.Find(&users)
	var userResult []models.UsersResult

	var GenZ, GenY, GenX, BabyBoomer, GI int = 0, 0, 0, 0, 0
	for _, user := range users {
		Genertion := ""
		if user.Age < 24 {
			GenZ++
			Genertion = "GenZ"
		} else if user.Age >= 24 && user.Age <= 41 {
			GenY++
			Genertion = "GenY"
		} else if user.Age >= 42 && user.Age <= 56 {
			GenX++
			Genertion = "GenX"
		} else if user.Age >= 57 && user.Age <= 75 {
			BabyBoomer++
			Genertion = "BabyBoomer"
		} else if user.Age > 75 {
			GI++
			Genertion = "G.I. Generation"
		}
		result := models.UsersResult{
			Name:      user.Name,
			Lastname:  user.Lastname,
			Birthday:  user.Birthday,
			Age:       user.Age,
			Genertion: Genertion,
		}
		userResult = append(userResult, result)
	}
	usersReturn := models.UsersReturn{
		Users:          userResult,
		GenZ:           GenZ,
		GenY:           GenY,
		GenX:           GenX,
		BabyBoomer:     BabyBoomer,
		G_I_Generation: GI,
	}

	return c.Status(200).JSON(usersReturn)
}

func UpdateUsers(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).JSON(err.Error())
	}
	db.Where("id = ?", id).Updates(&user)

	return c.Status(200).JSON(user)
}

func RemoveUsers(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var user models.Users
	db.Delete(&user, id)

	return c.Status(200).JSON(fiber.Map{
		"message": "Delete Success.",
		"status":  "ok",
	})
}
