package domain

import (
	"time"
)

type Seat struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	EventID uint `gorm:"not null" json:"event_id"` // FK ไปหา Events

	// ตรงนี้ map ชื่อ column ให้ตรงกับ SQL เป๊ะๆ
	RowLabel   string `gorm:"column:row_label;not null;uniqueIndex:idx_event_seat" json:"row_label"`
	SeatNumber int    `gorm:"column:seat_number;not null;uniqueIndex:idx_event_seat" json:"seat_number"`

	// หมายเหตุ: uniqueIndex:idx_event_seat ที่ซ้ำกัน 3 บรรทัด
	// คือการบอก GORM ว่าให้ทำ "UNIQUE(event_id, row_label, seat_number)"

	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`

	// ใช้ type:seat_status ที่เราจะสร้างเอง
	Status string `gorm:"column:status;type:seat_status;default:'AVAILABLE'" json:"status"`

	// FK ไปหา Users (Pointer เพราะอาจจะยังไม่มีคนจอง = NULL)
	LockedBy      *uint      `gorm:"column:locked_by" json:"locked_by"`
	LockExpiresAt *time.Time `gorm:"column:lock_expires_at" json:"lock_expires_at"`

	// Relation (Optional)
	Event *Event `gorm:"foreignKey:EventID" json:"-"`
}
