package main

import (
	"AdcashTask/Database"
	"AdcashTask/Models"
	"fmt"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	//category routes
	app.Get("/api/v1/Category", Models.GetCategories)
	app.Get("/api/v1/Category/:id", Models.GetCategory)
	app.Post("/api/v1/Category", Models.AddCategory)
	app.Delete("/api/v1/Category/:id", Models.DeleteCategories)
	app.Put("/api/v1/Category/:id", Models.UpdateCategories)

	//product routes
	app.Get("/api/v1/Products", Models.GetProducts)
	app.Get("/api/v1/Product/:id", Models.GetProduct)
	app.Post("/api/vgit1/Product", Models.AddProduct)
	app.Delete("/api/v1/Product/:id", Models.DeleteProduct)
	app.Put("/api/v1/Product/:id", Models.UpdateProduct)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")


	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection Successfully Opened ")
	database.DBConn.AutoMigrate(&Models.Category{}, &Models.Product{})


	fmt.Println("Database Migrated")
}

func main() {

	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Listen(":8080")

}
