package routes

import (
	employee_controller "tasteplorer-internal-api/app/controller"
	jwt_middleware "tasteplorer-internal-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Route for employee login (JWT generation)
	app.Post("/login", employee_controller.LoginEmployeeController)

	// Protected employee routes (JWT required)
	employee := app.Group("/employee") // Protect these routes with JWT middleware

	// CRUD operations for employees
	employee.Post("/", employee_controller.CreateEmployeeController)                              // Create new employee
	employee.Get("/:id", jwt_middleware.JWTMiddleware, employee_controller.GetEmployeeController) // Get employee by ID
	// employee.Put("/:id", controller.UpdateEmployeeController)    // Update employee
	// employee.Delete("/:id", controller.DeleteEmployeeController) // Delete employee
}
