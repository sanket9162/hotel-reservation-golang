package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "sanket",
		LastName:  "Gondhali",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("gondhali")
}
