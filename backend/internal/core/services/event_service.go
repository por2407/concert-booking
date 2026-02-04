package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ticket-backend/internal/core/domain"
	"github.com/ticket-backend/internal/core/ports"
)

type EventService struct {
	txManager ports.TransactionManager
	eventRepo ports.EventRepository
	seatRepo  ports.SeatRepository
}

func NewEventService(txManager ports.TransactionManager, evtRepo ports.EventRepository, seatRepo ports.SeatRepository) *EventService {
	return &EventService{
		txManager: txManager,
		eventRepo: evtRepo,
		seatRepo:  seatRepo,
	}
}

// ฟังก์ชันสร้างกิจกรรมพร้อมที่นั่ง
func (s *EventService) CreateEvent(ctx context.Context, name string, location string) (*domain.Event, error) {
	// 1. เตรียมข้อมูล Event
	event := domain.Event{
		Name:     name,
		Location: location,
		DateTime: time.Now().Add(30 * 24 * time.Hour),
	}

	// 2. ใช้ Transaction ครอบคลุมการสร้าง Event และ Seats
	err := s.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		// สร้าง Event ก่อน
		if err := s.eventRepo.Create(txCtx, &event); err != nil {
			return err
		}

		// สร้างที่นั่ง 50 ที่
		for i := 1; i <= 50; i++ {
			seat := domain.Seat{
				EventID:    event.ID, // ID จะถูกเติมให้หลังจาก Create สำเร็จ
				RowLabel:   "A",
				SeatNumber: i,
				Price:      1000,
				Status:     "AVAILABLE",
			}
			if err := s.seatRepo.Create(txCtx, &seat); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("🎉 Event '%s' and 50 seats created successfully!\n", event.Name)
	return &event, nil
}

func (s *EventService) GetEventInfo(ctx context.Context, eventID uint) (*domain.Event, error) {
	return s.eventRepo.GetByIDWithSeats(ctx, eventID)
}

func (s *EventService) GetAllEvents(ctx context.Context) ([]domain.Event, error) {
	return s.eventRepo.GetAll(ctx)
}
