package ports

import (
	"context"

	"github.com/ticket-backend/internal/core/domain"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error) // Returns JWT token
}

type BookingService interface {
	CreatePendingBooking(ctx context.Context, userID uint, seatIDs []uint) (*domain.Booking, error)
	ConfirmPayment(ctx context.Context, bookingID uint) (*domain.Booking, error)
	GetHistory(ctx context.Context, userID uint) ([]domain.Booking, error)
}

type EventService interface {
	CreateEvent(ctx context.Context, name string, location string) (*domain.Event, error)
	GetEventInfo(ctx context.Context, id uint) (*domain.Event, error)
	GetAllEvents(ctx context.Context) ([]domain.Event, error)
}
