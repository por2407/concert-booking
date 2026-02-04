package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ticket-backend/internal/core/ports"
)

type EventHandler struct {
	eventService ports.EventService
}

type createEventRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
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

	// 2. เรียก Service (ผ่าน UserContext ของ Fiber)
	event, err := h.eventService.GetEventInfo(c.UserContext(), uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
	}

	// 3. ส่ง JSON กลับไป
	return c.JSON(event)
}

func (h *EventHandler) GetAllEvents(c *fiber.Ctx) error {
	events, err := h.eventService.GetAllEvents(c.UserContext())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get events"})
	}
	return c.JSON(events)
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var req createEventRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	event, err := h.eventService.CreateEvent(c.UserContext(), req.Name, req.Location)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event created successfully",
		"event":   event,
	})
}
