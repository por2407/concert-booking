package services

import (
	"fmt"
	"time"

	"github.com/ticket-backend/internal/core/domain"
	"github.com/ticket-backend/internal/core/ports"
)

type eventService struct {
	eventRepo ports.EventRepository
	seatRepo  ports.SeatRepository
}

func NewEventService(evtRepo ports.EventRepository, seatRepo ports.SeatRepository) *eventService {
	return &eventService{
		eventRepo: evtRepo,
		seatRepo:  seatRepo,
	}
}

// ฟังก์ชันเสกข้อมูล
func (s *eventService) SeedData() error {
	// 1. สร้างคอนเสิร์ต
	event := domain.Event{
		Name:     "Microservice Concert",
		Location: "Bangkok Hall",
		DateTime: time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.eventRepo.Create(&event); err != nil {
		return err
	}
	fmt.Println("🎉 Event created:", event.Name)

	// 2. เสกที่นั่ง 50 ที่
	for i := 1; i <= 50; i++ {
		seat := domain.Seat{
			EventID:    event.ID,
			RowLabel:   "A",
			SeatNumber: i,
			Price:      1000,
			Status:     "AVAILABLE",
		}
		if err := s.seatRepo.Create(&seat); err != nil {
			return err
		}
	}
	fmt.Println("✅ 50 Seats generated!")
	return nil
}
