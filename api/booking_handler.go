package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotFount("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return err
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	if err := h.store.Booking.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return nil
	}
	return c.JSON(genericResp{
		Type: "msg",
		Msg:  "update",
	})
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	booking, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return ErrNotFount("booking")
	}
	return c.JSON(booking)

}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotFount("Booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}
