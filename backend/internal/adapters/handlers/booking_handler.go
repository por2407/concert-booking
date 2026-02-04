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

	userId := c.Locals("user_id").(uint)
	booking, err := h.bookingService.CreatePendingBooking(c.UserContext(), userId, req.SeatIds)
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

	booking, err := h.bookingService.ConfirmPayment(c.UserContext(), req.BookingId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Payment confirmed, seats are now SOLD",
		"booking": booking,
	})
}

func (h *BookingHandler) GetHistory(c *fiber.Ctx) error {
	// 1. ดึง User ID ออกมาจาก Context (ที่ถูกเซ็ตไว้โดย JWT Middleware)
	userID := c.Locals("user_id").(uint)
	bookingItems, err := h.bookingService.GetHistory(c.UserContext(), userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"booking_items": bookingItems})
}
