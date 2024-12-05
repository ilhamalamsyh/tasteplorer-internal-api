package upload_controller

import (
	"net/http"
	response_dto "tasteplorer-internal-api/app/dto/response"
	upload_dto "tasteplorer-internal-api/app/dto/upload"
	utils_upload_file "tasteplorer-internal-api/pkg/utils/upload"

	"github.com/gofiber/fiber/v2"
)

func UploadSingleFileController(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto

	uploadDto := new(upload_dto.UploadFileDto)

	if err := c.BodyParser(uploadDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request data",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		responseDto = response_dto.ResponseDto{
			Message: "Bad Request",
			Code:    http.StatusBadRequest,
			Error: fiber.Map{
				"message": "File is required",
			},
			Data: nil,
		}

		return c.Status(http.StatusBadRequest).JSON(responseDto)
	}

	uploadDto.File = file

	// Validate the DTO
	if err := uploadDto.Validate(); err != nil {
		responseDto = response_dto.ResponseDto{
			Message: "Bad Request",
			Code:    http.StatusBadRequest,
			Error: fiber.Map{
				"message": err.Error(),
			},
			Data: nil,
		}

		return c.Status(http.StatusBadRequest).JSON(responseDto)
	}

	// Use the folder from the DTO for Cloudinary upload
	uploadResult, err := utils_upload_file.UploadToCloudinary(c.Context(), uploadDto.File, string(uploadDto.Path))
	if err != nil {
		responseDto = response_dto.ResponseDto{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
			Error: fiber.Map{
				"message": err.Error(),
			},
			Data: nil,
		}

		return c.Status(http.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "OK",
		Code:    http.StatusOK,
		Error:   nil,
		Data: fiber.Map{
			"message":   "File uploaded successfully",
			"url":       uploadResult.SecureURL,
			"folder":    uploadDto.Path,
			"public_id": uploadResult.PublicID,
		},
	}

	return c.Status(http.StatusOK).JSON(responseDto)
}
