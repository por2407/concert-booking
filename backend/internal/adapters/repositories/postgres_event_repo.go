package repositories

import (
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresEventRepo struct {
	db *gorm.DB
}

func NewPostgresEventRepo(db *gorm.DB) *PostgresEventRepo {
	return &PostgresEventRepo{db: db}
}

func (r *PostgresEventRepo) Create(event *domain.Event) error {
	return r.db.Create(event).Error
}

func (r *PostgresEventRepo) GetByIDWithSeats(id uint) (*domain.Event, error) {
	var event domain.Event
	// Preload("Seats") คือบอก GORM ว่า "ไปดึงตาราง seats ที่เกี่ยวข้องมาใส่ในฟิลด์ Seats ด้วยนะ"
	// เรียงลำดับตาม Row และ Number เพื่อให้ Frontend วาดง่ายๆ
	err := r.db.Preload("Seats", func(db *gorm.DB) *gorm.DB {
		return db.Order("row_label ASC, seat_number ASC")
	}).First(&event, id).Error

	return &event, err
}
