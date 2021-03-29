package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jnjam6681/fiber-gorm/book"
	"github.com/jnjam6681/fiber-gorm/database"
)

func setupRouter(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error

	// dsn := "root:root@tcp(127.0.0.1:3306)/fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	database.DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open(mysql.Open("books.db"), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Database successfully opened")

	database.DB.AutoMigrate(&book.Book{})
	fmt.Println("Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()

	defer database.DB.Close()

	// app.Get("/", helloworld)
	setupRouter(app)

	app.Listen(":3000")
}

// func helloworld(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World")
// }
