package routes

import (
	"github.com/Hammarang/WhereAreWeGoing/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(app *fiber.App) {
	gameRouter := app.Group("/games")
	gameRouter.Get("/", controllers.GetAllGames)
}
