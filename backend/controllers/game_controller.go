package controllers

import (
	"fmt"

	"github.com/Hammarang/WhereAreWeGoing/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllGames(c *fiber.Ctx) error {
	games := models.GetAllGames()
	fmt.Println(games)
	return c.JSON(games)
}
