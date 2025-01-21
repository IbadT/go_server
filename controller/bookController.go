package controller

import (
	"net/http"

	"github.com/IbadT/go_server/config"
	. "github.com/IbadT/go_server/model"
	"github.com/labstack/echo/v4"
)

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} Book
// @Router /book/ [get]
func GetAllBooks(c echo.Context) error {
	db := config.DB()
	var books []*Book
	if err := db.Find(&books).Error; err != nil {
		// if err := db.Select("id", "name", "description", "created_at", "updated_at").Find(&books).Error; err != nil {
		// if err := db.Preload("Author").Preload("Publisher").Preload("Categories").Find(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Книги не найдены"})
	}
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")

	db := config.DB()
	var book *Book
	if err := db.Preload("Author").Preload("Publisher").Preload("Categories").First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := config.DB()
	if err := db.Create(book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create book"})
	}

	return c.JSON(http.StatusCreated, book)
}

// func GetBook(c echo.Context) error {
// 	id := c.Param("id")
// 	db := config.DB()

// 	var books []*model.Book

// 	if err := db.Find(&books, id).Error; err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusOK, data)
// 	}

// 	response := map[string]interface{}{
// 		"data": books[0],
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func GetAllBooks(c echo.Context) error {
// 	db := config.DB()

// 	var books []*model.Book
// 	if err := db.Find(&books).Error; err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusOK, data)
// 	}

// 	response := map[string]interface{}{
// 		"data": books,
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func CreateBook(c echo.Context) error {
// 	b := new(model.Book)
// 	db := config.DB()

// 	// Binding data
// 	if err := c.Bind(b); err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusInternalServerError, data)
// 	}

// 	book := &model.Book{
// 		Name:        b.Name,
// 		Description: b.Description,
// 	}

// 	if err := db.Create(&book).Error; err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusInternalServerError, data)
// 	}

// 	response := map[string]interface{}{
// 		"data": b,
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func UpdateBook(c echo.Context) error {
// 	id := c.Param("id")
// 	b := new(model.Book)
// 	db := config.DB()

// 	// Binding data
// 	if err := c.Bind(b); err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusInternalServerError, data)
// 	}

// 	existing_book := new(model.Book)

// 	if err := db.First(&existing_book, id).Error; err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusNotFound, data)
// 	}

// 	existing_book.Name = b.Name
// 	existing_book.Description = b.Description
// 	if err := db.Save(&existing_book).Error; err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusInternalServerError, data)
// 	}

// 	response := map[string]interface{}{
// 		"data": existing_book,
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func DeleteBook(c echo.Context) error {
// 	id := c.Param("id")
// 	db := config.DB()

// 	book := new(model.Book)

// 	if err := db.Delete(&book, id).Error; err != nil {
// 		// err := db.Delete(&book, id).Error
// 		// if err != nil {
// 		data := map[string]interface{}{
// 			"message": err.Error(),
// 		}

// 		return c.JSON(http.StatusInternalServerError, data)
// 	}

// 	response := map[string]interface{}{
// 		"message": "a book has been deleted",
// 	}
// 	return c.JSON(http.StatusOK, response)
// }
