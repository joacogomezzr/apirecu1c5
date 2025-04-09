// api_hexagonal_go/internal/book/interface/book_handler.go
package interfaces

import (
	"api-joaquin/internal/book/controllers"
	"api-joaquin/internal/book/domain"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	Controller *controllers.BookController
}

func NewBookHandler(controller *controllers.BookController) *BookHandler {
	return &BookHandler{Controller: controller}
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	type BookRequest struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	var req BookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al parsear los datos del libro",
			"error":   err.Error(),
		})
	}

	if req.Title == "" || req.Author == "" || req.Year == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Todos los campos son requeridos (título, autor, año)",
		})
	}

	currentYear := time.Now().Year()
	if req.Year < 1000 || req.Year > currentYear {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("El año debe estar entre 1000 y %d", currentYear),
		})
	}

	book := &domain.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := h.Controller.PostUseCase.Execute(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error al guardar el libro en la base de datos",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro creado exitosamente",
		"data": fiber.Map{
			"title":  book.Title,
			"author": book.Author,
			"year":   book.Year,
		},
	})
}

// GetBooks maneja la obtención de libros.
func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.Controller.GetUseCase.Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudieron obtener los libros.",
			"error":   err.Error(),
		})
	}

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "success",
			"message": "No se encontraron libros en la base de datos.",
			"data":    books,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Lista de libros obtenida con éxito.",
		"data":    books,
	})
}
// UpdateBook maneja la actualización de un libro.
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido.",
		})
	}

	book := new(domain.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Solicitud inválida. Verifique los datos enviados.",
		})
	}

	book.ID = id
	if err := h.Controller.PutUseCase.Execute(book); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Libro no encontrado.",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo actualizar el libro.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro actualizado correctamente.",
		"data":    book,
	})
}

// DeleteBook maneja la eliminación de un libro.
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido.",
		})
	}

	if err := h.Controller.DeleteUseCase.Execute(id); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Libro no encontrado.",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo eliminar el libro.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro eliminado correctamente.",
	})
}
