package ports

import (
	"github.com/ticket-backend/internal/core/domain"
)

// กำหนดว่า คนที่จะมาจัดการ Event ต้องทำสิ่งนี้ได้นะ
type EventRepository interface {
	Create(event *domain.Event) error
}

// กำหนดว่า คนที่จะมาจัดการ Seat ต้องทำสิ่งนี้ได้นะ
type SeatRepository interface {
	Create(seat *domain.Seat) error
}
