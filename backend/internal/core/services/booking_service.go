package services

import (
	"github.com/ticket-backend/internal/adapters/repositories"
	"gorm.io/gorm"
)

type bookingService struct {
	db          *gorm.DB // เราต้องใช้ DB เพื่อสั่ง Start Transaction
	seatRepo    *repositories.PostgresSeatRepo
	bookingRepo *repositories.PostgresBookingRepo
}

// รับ Repo เข้ามา (สังเกตว่าผมใช้ Type จริง (struct) แทน Interface ชั่วคราวเพื่อให้ง่ายต่อการส่ง Transaction)
func NewBookingService(db *gorm.DB, seatRepo *repositories.PostgresSeatRepo, bookingRepo *repositories.PostgresBookingRepo) *bookingService {
	return &bookingService{
		db:          db,
		seatRepo:    seatRepo,
		bookingRepo: bookingRepo,
	}
}
