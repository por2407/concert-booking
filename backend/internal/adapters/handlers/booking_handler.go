package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ticket-backend/internal/core/ports"
)

type BookingHandler struct {
	bookingService ports.BookingService
}

func NewBookingHandler(s ports.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: s}
}

type createBookingRequest struct {
	UserId  uint   `json:"user_id"`
	SeatIds []uint `json:"seat_ids"`
}

type confirmBookingRequest struct {
	BookingId uint `json:"booking_id"`
}

// 1. กดจอง (รอชำระเงิน)
func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var req createBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	booking, err := h.bookingService.CreatePendingBooking(req.UserId, req.SeatIds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Booking pending, please pay within 10 minutes",
		"booking": booking,
	})
}

// 2. ยืนยันการชำระเงิน (Sold)
func (h *BookingHandler) ConfirmBooking(c *fiber.Ctx) error {
	var req confirmBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	booking, err := h.bookingService.ConfirmPayment(req.BookingId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Payment confirmed, seats are now SOLD",
		"booking": booking,
	})
}
