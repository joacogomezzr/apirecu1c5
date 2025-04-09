package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"api-joaquin/config"
	"api-joaquin/database"
	"api-joaquin/pkg/middleware"


	bookControllers "api-joaquin/internal/book/controllers"
	bookInfrastructure "api-joaquin/internal/book/infrastructure"
	bookInterfaces "api-joaquin/internal/book/interface"
	bookRoutes "api-joaquin/internal/book/routes"

	adminControllers "api-joaquin/internal/admin/controllers"
	adminInfrastructure "api-joaquin/internal/admin/infrastructure"
	adminInterfaces "api-joaquin/internal/admin/interface"
	adminRoutes "api-joaquin/internal/admin/routes"
)

func main() {
	config.LoadEnv()

	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if config.GetEnv(v, "") == "" {
			log.Fatalf("❌ Error: La variable de entorno %s no está configurada", v)
		}
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "API Biblioteca",
		AppName:      "Biblioteca v1.0",
	})

	app.Use(middleware.SetupCORS())


	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${status} - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     os.Stdout,
	}))

	bookRepo := bookInfrastructure.NewBookService(database.DB)
	bookController := bookControllers.NewBookController(bookRepo)
	bookHandler := bookInterfaces.NewBookHandler(bookController)

	adminRepo := adminInfrastructure.NewAdminService(database.DB)
	adminController := adminControllers.NewAdminController(adminRepo)
	adminHandler := adminInterfaces.NewAdminHandler(adminController)

	bookRoutes.SetupBookRoutes(app, bookHandler)
	adminRoutes.SetupAdminRoutes(app, adminHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Endpoint no encontrado",
		})
	})

	port := config.GetEnv("PORT", "8080")
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}