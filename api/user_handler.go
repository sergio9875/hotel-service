package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sergio9875/hotel-service/db"
	"github.com/sergio9875/hotel-service/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		ctx = context.Background()
		id  = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Sergey",
		LastName:  "Styles",
	}
	return c.JSON(u)
}
