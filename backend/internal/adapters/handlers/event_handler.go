package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ticket-backend/internal/core/ports"
)

type EventHandler struct {
	eventService ports.EventService
}

func NewEventHandler(s ports.EventService) *EventHandler {
	return &EventHandler{eventService: s}
}

func (h *EventHandler) GetEvent(c *fiber.Ctx) error {
	// 1. อ่านค่า id จาก URL (เช่น /api/events/1)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid event ID"})
	}

	// 2. เรียก Service
	event, err := h.eventService.GetEventInfo(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
	}

	// 3. ส่ง JSON กลับไป
	return c.JSON(event)
}
