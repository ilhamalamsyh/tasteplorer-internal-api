package article_controller

import (
	"log"
	"math"
	"strconv"
	article_dto "tasteplorer-internal-api/app/dto/article"
	metadata_dto "tasteplorer-internal-api/app/dto/meta"
	response_dto "tasteplorer-internal-api/app/dto/response"
	article_service "tasteplorer-internal-api/app/service/article"
	utils_validation "tasteplorer-internal-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateArticleController(c *fiber.Ctx) error {
	var articleRequestDto article_dto.ArticleRequestDto
	var responseDto response_dto.ResponseDto

	if err := c.BodyParser(&articleRequestDto); err != nil {
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

	customMessages := articleRequestDto.CustomMessagesValidation()
	validationErrors := utils_validation.ValidateStruct(articleRequestDto, customMessages)
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

	article, err := article_service.CreateArticleService(&articleRequestDto)

	if err != nil {
		log.Printf("error creating new article: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error when creating new article",
			Code:    fiber.StatusBadRequest,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success create new article",
		Code:    fiber.StatusCreated,
		Error:   nil,
		Data:    article,
	}
	return c.Status(fiber.StatusCreated).JSON(responseDto)
}

func GetAllArticleController(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto
	var metaDataDto metadata_dto.MetaData

	page := c.Query("page", "0")
	pageSize := c.Query("pageSize", "10")
	search := c.Query("search", "")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 0 {
		pageInt = 0
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 1
	}

	articles, total, err := article_service.FindAllArticleService(uint(pageInt), uint(pageSizeInt), search)
	if err != nil {
		log.Printf("error fetching article list: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error fetching article list",
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
			"articles": articles,
			"meta":     metaDataDto,
		},
	}

	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func GetArticleDetailController(c *fiber.Ctx) error {
	id := c.Params("id")
	var responseDto response_dto.ResponseDto

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	article, err := article_service.ArticleDetailService(uint(idInt))

	if article == nil {
		responseDto = response_dto.ResponseDto{
			Message: "Article not found",
			Code:    fiber.StatusNotFound,
			Error: fiber.Map{
				"message": "Article not found.	",
			},
			Data: nil,
		}

		return c.Status(fiber.StatusNotFound).JSON(responseDto)
	}

	if err != nil {
		log.Printf("Error get article detail: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error get article detail",
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
		Data:    article,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func UpdateArticleController(c *fiber.Ctx) error {
	id := c.Params("id")
	var articleRequestDto article_dto.ArticleRequestDto
	var responseDto response_dto.ResponseDto

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&articleRequestDto); err != nil {
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

	article, err := article_service.UpdateArticleService(uint(idInt), &articleRequestDto)

	if err != nil {
		log.Printf("Error when updating article: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error update article",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err,
			},
			Data: nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success update article",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    article,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)
}

func DeleteArticleContoller(c *fiber.Ctx) error {
	var responseDto response_dto.ResponseDto
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = article_service.DeleteArticleService(uint(idInt))

	if err != nil {
		log.Printf("Error when deleting article: %v", err)
		responseDto = response_dto.ResponseDto{
			Message: "Error when deleting article",
			Code:    fiber.StatusInternalServerError,
			Error: fiber.Map{
				"message": err.Error(),
			},
			Data: nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(responseDto)
	}

	responseDto = response_dto.ResponseDto{
		Message: "Success delete article",
		Code:    fiber.StatusOK,
		Error:   nil,
		Data:    nil,
	}
	return c.Status(fiber.StatusOK).JSON(responseDto)

}
