package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sergio9875/hotel-service/types"
)

func HandleUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Sergey",
		LastName:  "Styles",
	}
	return c.JSON(u)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("User by id is:")

}
