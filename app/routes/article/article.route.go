package article_router

import (
	article_controller "tasteplorer-internal-api/app/controller/article"
	jwt_middleware "tasteplorer-internal-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	article := app.Group("/articles", jwt_middleware.JWTMiddleware)

	article.Get("/", article_controller.GetAllArticleController)
	article.Get("/:id", article_controller.GetArticleDetailController)
	article.Post("/", article_controller.CreateArticleController)
	article.Put("/:id", article_controller.UpdateArticleController)
	article.Delete("/:id", article_controller.DeleteArticleContoller)

}
