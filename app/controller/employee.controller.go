package employee_controller

import (
	"fmt"
	"log"
	"strconv"
	employee_dto "tasteplorer-internal-api/app/dto/employee"
	response_dto "tasteplorer-internal-api/app/dto/response"
	employee_service "tasteplorer-internal-api/app/service/employee"
	utils_validation "tasteplorer-internal-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginEmployeeController(c *fiber.Ctx) error {
	var loginDto employee_dto.LoginRequest
	var responseDto response_dto.ResponseDto

	if err := c.BodyParser(&loginDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Payload")
	}

	customMessages := loginDto.CustomMessagesValidationX()
	validationErrors := utils_validation.ValidateStruct(loginDto, customMessages)

	if len(validationErrors) > 0 {
		fmt.Println("eyoy: ", validationErrors)
		var errorMessages []string
		for _, message := range validationErrors {
			errorMessages = append(errorMessages, message)
		}

		responseDto = response_dto.ResponseDto{
			Message: "Invalid request payload",
			Code:    fiber.StatusUnprocessableEntity,
			Error: fiber.Map{
				"message": errorMessages[0],
			},
			Data: nil,
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(responseDto)
	}

	token, err := employee_service.LoginService(&loginDto)
	if err != nil {
		log.Printf("Error logging in employee: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Bad Request",
			Code:    fiber.StatusBadRequest,
			Error: fiber.Map{
				"message": err.Error(),
			},
			Data: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Login berhasil",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data: fiber.Map{
			"token": token,
		},
	}

	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func CreateEmployeeController(c *fiber.Ctx) error {
	var registerDto employee_dto.RegisterDto
	var responseDto response_dto.ResponseDto

	if err := c.BodyParser(&registerDto); err != nil {
		responseDto = response_dto.ResponseDto{
			Message: "Invalid request payload",
			Code:    fiber.StatusBadRequest,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(responseDto)
	}

	employee, err := employee_service.CreateEmployeeService(&registerDto)

	if err != nil {
		log.Printf("Error creating employee: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success register new employee",
		Code:    fiber.StatusCreated,
		Error:   nil,
		Data:    employee,
	}
	return c.Status(fiber.StatusCreated).JSON(responseDto)
}

func GetEmployeeController(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto

	id := c.Params("id")

	convertedId, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		log.Println("Error:", err) // Output: Error: strconv.ParseUint: parsing "abc123": invalid syntax
	}

	employee, err := employee_service.EmployeeDetailService(uint(convertedId))
	if err != nil {
		log.Printf("Error retrieving employee: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Employee not found",
			Code:    fiber.StatusNotFound,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusNotFound).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success get employee detail",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    employee,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)
}
