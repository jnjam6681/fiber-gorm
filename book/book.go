package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jnjam6681/fiber-gorm/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.DB
	// var book Book
	// book.Title = "Hi"
	// book.Author = "JJ"
	// book.Rating = 5

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	db.Create(&book)
	return c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
