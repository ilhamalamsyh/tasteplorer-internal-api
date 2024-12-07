package upload_router

import (
	upload_controller "tasteplorer-internal-api/app/controller/upload"
	jwt_middleware "tasteplorer-internal-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	upload := router.Group("/upload", jwt_middleware.JWTMiddleware)

	upload.Post("/", upload_controller.UploadSingleFileController)

}
