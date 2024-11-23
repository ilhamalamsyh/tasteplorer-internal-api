package employee_controller

import (
	"log"
	"strconv"
	employee_dto "tasteplorer-internal-api/app/dto/employee"
	dto_custom_error "tasteplorer-internal-api/app/dto/errors"

	employee_service "tasteplorer-internal-api/app/service/employee"

	"github.com/gofiber/fiber/v2"
)

func LoginEmployeeController(c *fiber.Ctx) error {
	var loginDto employee_dto.LoginRequest
	var customeError dto_custom_error.CustomError

	if err := c.BodyParser(&loginDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Payload")
	}

	token, err := employee_service.LoginService(&loginDto)
	if err != nil {
		log.Printf("Error logging in employee: %v", err)
		customeError = dto_custom_error.CustomError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(customeError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
	})
}

func CreateEmployeeController(c *fiber.Ctx) error {
	var registerDto employee_dto.RegisterDto
	var customeError dto_custom_error.CustomError

	if err := c.BodyParser(&registerDto); err != nil {
		customeError = dto_custom_error.CustomError{
			Message: "Invalid request payload",
			Code:    fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(customeError)
	}

	employee, err := employee_service.CreateEmployeeService(&registerDto)

	if err != nil {
		log.Printf("Error creating employee: %v", err)
		customeError = dto_custom_error.CustomError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}
		return c.Status(fiber.StatusBadRequest).JSON(customeError)
	}

	return c.Status(fiber.StatusCreated).JSON(employee)
}

func GetEmployeeController(c *fiber.Ctx) error {
	var customeError dto_custom_error.CustomError

	id := c.Params("id")

	convertedId, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		log.Println("Error:", err) // Output: Error: strconv.ParseUint: parsing "abc123": invalid syntax
	}

	employee, err := employee_service.EmployeeDetailService(uint(convertedId))
	if err != nil {
		log.Printf("Error retrieving employee: %v", err)
		customeError = dto_custom_error.CustomError{
			Message: "Employee not found",
			Code:    fiber.StatusNotFound,
		}
		return c.Status(fiber.StatusNotFound).JSON(customeError)
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}
