package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/karnpini/fiber-basic-crud/customer"
	"github.com/karnpini/fiber-basic-crud/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	db, err := gorm.Open(sqlite.Open("customer.db"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&customer.Customer{})
	database.DBConn = db
	fmt.Println("Database Migrated")
}

func setupRoute(app *fiber.App) {
	app.Get("/customers", customer.GetAllCustomer)
	app.Get("/customers/:id", customer.GetCustomer)
	app.Post("/customers", customer.SaveCustomer)
	app.Put("/customers/:id", customer.UpdateCustomer)
	app.Delete("/customers/:id", customer.DeleteCustomer)
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoute(app)

	app.Listen(":3001")
}
