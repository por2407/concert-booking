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

// โครงสร้าง Request ที่จะรับจาก Frontend
// หน้าตา JSON จะเป็น: { "user_id": 1, "seat_ids": [1, 2] }

type createBookingRequest struct {
	UserId  uint   `json:"user_id"`
	SeatIds []uint `json:"seat_ids"`
}

// ฟังก์ชันที่จะถูกเรียกเมื่อมีคนยิง API เข้ามา
func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	// 1. รับ JSON และแปลงเข้า Struct
	var req createBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	// 2. เรียกใช้ Service (Logic การจองที่เราทำไปแล้ว)
	booking, err := h.bookingService.CreateBooking(req.UserId, req.SeatIds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Booking created successfully", "booking": booking})

}
