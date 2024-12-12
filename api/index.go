package main

import (
	"log"
	"net/http"

	article_router "tasteplorer-internal-api/app/routes/article"
	banner_router "tasteplorer-internal-api/app/routes/banner"
	employee_router "tasteplorer-internal-api/app/routes/employee"
	upload_router "tasteplorer-internal-api/app/routes/upload"
	"tasteplorer-internal-api/platform/database"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// Initialize configurations and database connection
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	database.Initialize()
}

// Fiber application instance
var app = fiber.New()

// Register routes
func setupRoutes() {
	// Enable CORS
	app.Use(cors.New())

	// Define API group
	api := app.Group("/api")

	// Setup routes from different modules
	employee_router.SetupRoutes(api)
	banner_router.SetupRoutes(api)
	article_router.SetupRoutes(api)
	upload_router.SetupRoutes(api)
}

// Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	// Ensure routes are set up
	setupRoutes()

	// Use Fiber's adaptor to serve the request
	adaptor.FiberApp(app).ServeHTTP(w, r)
}

func main() {
	// For local development
	setupRoutes()
	log.Fatal(app.Listen(":5000"))
}
