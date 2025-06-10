package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/db"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	return nil
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	return nil
}
