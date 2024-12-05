package main

import (
	"log"
	article_router "tasteplorer-internal-api/app/routes/article"
	banner_router "tasteplorer-internal-api/app/routes/banner"
	routes "tasteplorer-internal-api/app/routes/employee"
	upload_router "tasteplorer-internal-api/app/routes/upload"
	"tasteplorer-internal-api/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database connection
	database.Initialize()
}

func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New())

	// Initialize Controllers
	routes.SetupRoutes(app)
	banner_router.SetupRoutes(app)
	article_router.SetupRoutes(app)
	upload_router.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":5000"))

}
