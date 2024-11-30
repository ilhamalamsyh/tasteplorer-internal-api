package banner_router

import (
	banner_controller "tasteplorer-internal-api/app/controller/banner"
	jwt_middleware "tasteplorer-internal-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	banner := app.Group("/banners", jwt_middleware.JWTMiddleware)

	banner.Get("/", banner_controller.GetAllBannerController)
	banner.Get("/:id", banner_controller.GetBannerDetailController)
	banner.Post("/", banner_controller.CreateBannerController)
	banner.Put("/:id", banner_controller.UpdateBannerController)
	banner.Delete("/:id", banner_controller.DeleteBannerContoller)

}
