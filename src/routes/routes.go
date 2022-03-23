package routes

import (
	"image-upload-go/src/controllers"
	// "complete_golang_backend/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	main := api.Group("main")
	main.Post("image-upload", controllers.UploadFile)

}