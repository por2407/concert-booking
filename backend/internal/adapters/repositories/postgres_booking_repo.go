package repositories

import (
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresBookingRepo struct {
	db *gorm.DB
}

func NewPostgresBookingRepo(db *gorm.DB) *PostgresBookingRepo {
	return &PostgresBookingRepo{db: db}
}

// ฟังก์ชันสร้างใบจอง (รองรับ Transaction)
func (r *PostgresBookingRepo) Create(tx *gorm.DB, booking *domain.Booking) error {
	// ถ้าส่ง Transaction มาให้ใช้ Transaction, ถ้าไม่มีให้ใช้ DB ปกติ
	if tx != nil {
		return tx.Create(booking).Error
	}
	return r.db.Create(booking).Error
}
