package customer

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/karnpini/fiber-basic-crud/database"
)

type Customer struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func GetAllCustomer(ctx *fiber.Ctx) error {
	db := database.DBConn
	customer := []Customer{}
	db.Find(&customer)
	return ctx.Status(http.StatusOK).JSON(customer)
}

func GetCustomer(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")
	customer := []Customer{}

	if err := db.Find(&customer, id).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(customer)
}

func SaveCustomer(ctx *fiber.Ctx) error {
	db := database.DBConn
	customer := new(Customer)

	if err := ctx.BodyParser(customer); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := db.Save(&customer).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(customer)
}

func UpdateCustomer(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")
	customer := new(Customer)

	if err := db.Find(&customer, id).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	if err := ctx.BodyParser(customer); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := db.Save(&customer).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(customer)
}

func DeleteCustomer(ctx *fiber.Ctx) error {
	db := database.DBConn
	id := ctx.Params("id")
	customer := new(Customer)

	if err := db.Find(&customer, id).Error; err != nil {
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	if err := db.Delete(&customer).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(customer)
}
