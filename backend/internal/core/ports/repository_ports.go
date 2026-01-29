package ports

import (
	"github.com/ticket-backend/internal/core/domain"
)

// กำหนดว่า คนที่จะมาจัดการ Event ต้องทำสิ่งนี้ได้นะ
type EventRepository interface {
	Create(event *domain.Event) error
	// เพิ่มฟังก์ชันนี้: ดึงงานตาม ID พร้อมข้อมูลที่นั่งทั้งหมด
	GetByIDWithSeats(id uint) (*domain.Event, error)
}

// กำหนดว่า คนที่จะมาจัดการ Seat ต้องทำสิ่งนี้ได้นะ
type SeatRepository interface {
	Create(seat *domain.Seat) error
}
