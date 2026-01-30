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
	if tx != nil {
		return tx.Create(booking).Error
	}
	return r.db.Create(booking).Error
}

func (r *PostgresBookingRepo) GetByID(id uint) (*domain.Booking, error) {
	var booking domain.Booking
	// Preload Items เพื่อให้รู้ว่าจองที่นั่งไหนบ้าง
	if err := r.db.Preload("Items").First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *PostgresBookingRepo) Update(tx *gorm.DB, booking *domain.Booking) error {
	if tx != nil {
		return tx.Save(booking).Error
	}
	return r.db.Save(booking).Error
}
