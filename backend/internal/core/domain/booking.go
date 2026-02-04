package domain

import "time"

// 1. สร้าง Enum Status (เดี๋ยวเราไปสร้าง Type ใน DB ทีหลังเหมือน seat_status)
type BookingStatus string

const (
	BookingPending   BookingStatus = "PENDING"
	BookingPaid      BookingStatus = "PAID"
	BookingCancelled BookingStatus = "CANCELLED"
)

// 2. ตาราง Bookings (ใบจองหลัก)
type Booking struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	UserID      uint          `gorm:"column:user_id;not null" json:"user_id"`
	EventID     uint          `gorm:"column:event_id;not null" json:"event_id"`
	TotalAmount float64       `gorm:"column:total_amount;type:decimal(10,2);not null" json:"total_amount"`
	Status      BookingStatus `gorm:"column:status;type:booking_status;default:'PENDING'" json:"status"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relation: ใบจอง 1 ใบ มีที่นั่งได้หลายที่ (Booking Items)
	Items []BookingItem `gorm:"foreignKey:BookingID" json:"items"`
	Event *Event        `gorm:"foreignKey:EventID" json:"event,omitempty"`
}

// 3. ตาราง BookingItems (รายละเอียดว่าจองที่นั่งไหนบ้าง)
type BookingItem struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	BookingID uint `gorm:"column:booking_id;not null" json:"booking_id"`
	SeatID    uint `gorm:"column:seat_id;not null" json:"seat_id"`

	// Relation (Optional): เพื่อให้ดึงข้อมูลที่นั่งได้ง่ายๆ
	Seat *Seat `gorm:"foreignKey:SeatID" json:"seat,omitempty"`
}
