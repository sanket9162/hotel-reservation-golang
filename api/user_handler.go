package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/db"
	"github.com/sanket9162/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "sanket",
		LastName:  "Gondhali",
	}
	return c.JSON(u)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUderById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
