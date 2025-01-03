package handler

import (
	"log"
	"net/http"
	"os"

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
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else {
		log.Println(".env file not found, using environment variables")
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

func handler() http.HandlerFunc {
	// Ensure routes are set up
	setupRoutes()

	return adaptor.FiberApp(app)
}

// Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()
	// Use Fiber's adaptor to serve the request
	handler().ServeHTTP(w, r)
}
