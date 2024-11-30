package banner_controller

import (
	"log"
	"math"
	"strconv"
	banner_dto "tasteplorer-internal-api/app/dto/banner"
	metadata_dto "tasteplorer-internal-api/app/dto/meta"
	response_dto "tasteplorer-internal-api/app/dto/response"
	banner_service "tasteplorer-internal-api/app/service/banner"
	utils_validation "tasteplorer-internal-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateBannerController(c *fiber.Ctx) error {
	var bannerRequestDto banner_dto.BannerRequestDto
	var responseDto response_dto.ResponseDto

	if err := c.BodyParser(&bannerRequestDto); err != nil {
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

	customMessages := bannerRequestDto.CustomMessagesValidation()
	validationErrors := utils_validation.ValidateStruct(bannerRequestDto, customMessages)
	if len(validationErrors) > 0 {
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

	banner, err := banner_service.CreateBannerService(&bannerRequestDto)

	if err != nil {
		log.Printf("error creating new banner: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error when creating new banner",
			Code:    fiber.StatusBadRequest,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success create new banner",
		Code:    fiber.StatusCreated,
		Error:   nil,
		Data:    banner,
	}
	return c.Status(fiber.StatusCreated).JSON(responseDto)
}

func GetAllBannerController(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto
	var metaDataDto metadata_dto.MetaData

	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")
	search := c.Query("search", "")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 1
	}

	banners, total, err := banner_service.FindAllBannerService(uint(pageInt), uint(pageSizeInt), search)
	if err != nil {
		log.Printf("error fetching banner list: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error fetching banner list",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}

		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	totalPage := int(math.Ceil(float64(total) / float64(pageSizeInt)))

	metaDataDto = metadata_dto.MetaData{
		CurrentPage: pageInt,
		PageSize:    pageSizeInt,
		Total:       total,
		TotalPage:   totalPage,
	}
	responseDto = response_dto.ResponseDto{
		Message: "Success fetching banner list",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data: fiber.Map{
			"banners": banners,
			"meta":    metaDataDto,
		},
	}

	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func GetBannerDetailController(c *fiber.Ctx) error {
	id := c.Params("id")
	var responseDto response_dto.ResponseDto

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	banner, err := banner_service.BannerDetailService(uint(idInt))

	if banner == nil {
		responseDto = response_dto.ResponseDto{
			Message: "Banner not found",
			Code:    fiber.StatusNotFound,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}

		return c.Status(fiber.StatusNotFound).JSON(responseDto)
	}

	if err != nil {
		log.Printf("Error get banner detail: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error get banner detail",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}

		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Sucess get banner detail",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    banner,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func UpdateBannerController(c *fiber.Ctx) error {
	id := c.Params("id")
	var bannerRequestDto banner_dto.BannerRequestDto
	var responseDto response_dto.ResponseDto

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&bannerRequestDto); err != nil {
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

	banner, err := banner_service.UpdateBannerService(uint(idInt), &bannerRequestDto)

	if err != nil {
		log.Printf("Error when updating banner: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error update banner",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success update banner",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    banner,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func DeleteBannerContoller(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = banner_service.DeleteBannerService(uint(idInt))

	if err != nil {
		log.Printf("Error when deleting banner: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error when deleting banner",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err.Error(),
			},
			Data: nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success delete banner",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    nil,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)

}
