package ports

import "github.com/ticket-backend/internal/core/domain"

// BookingService คือข้อตกลงว่า "ระบบจอง" ต้องทำอะไรได้บ้าง
// (ในที่นี้คือ ต้องมีฟังก์ชัน CreateBooking)
type BookingService interface {
	CreatePendingBooking(userID uint, seatIDs []uint) (*domain.Booking, error)
	ConfirmPayment(bookingID uint) (*domain.Booking, error)
}

type EventService interface {
	SeedData() error
	// เพิ่มฟังก์ชันดึงข้อมูล
	GetEventInfo(id uint) (*domain.Event, error)
}
