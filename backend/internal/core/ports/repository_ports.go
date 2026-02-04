package ports

import (
	"context"

	"github.com/ticket-backend/internal/core/domain"
)

// TransactionManager จัดการเรื่อง Transaction โดยไม่ให้ Service รู้จัก GORM
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
}

type EventRepository interface {
	Create(ctx context.Context, event *domain.Event) error
	GetByIDWithSeats(ctx context.Context, id uint) (*domain.Event, error)
	GetAll(ctx context.Context) ([]domain.Event, error)
}

type SeatRepository interface {
	Create(ctx context.Context, seat *domain.Seat) error
	FindByID(ctx context.Context, id uint) (*domain.Seat, error)
	Update(ctx context.Context, seat *domain.Seat) error
	GetSeatWithLock(ctx context.Context, seatID uint) (*domain.Seat, error)
}

type BookingRepository interface {
	Create(ctx context.Context, booking *domain.Booking) error
	GetByID(ctx context.Context, id uint) (*domain.Booking, error)
	Update(ctx context.Context, booking *domain.Booking) error
	GetByUserID(ctx context.Context, userID uint) ([]domain.Booking, error)
}
