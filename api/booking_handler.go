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

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	booking, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(booking)

}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	return nil
}
