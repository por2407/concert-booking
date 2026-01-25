package domain

import "time"

type Event struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	// ใช้ column:date_time
	DateTime time.Time `gorm:"column:date_time;not null" json:"date_time"`
	Location string    `gorm:"not null" json:"location"`
	IsActive bool      `gorm:"default:true" json:"is_active"`

	// Relation: เอาไว้จอยตาราง (GORM จะไม่สร้าง column นี้ใน db แต่ใช้ตอน query)
	Seats []Seat `gorm:"foreignKey:EventID" json:"seats,omitempty"`
}
