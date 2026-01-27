package repositories

import (
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresSeatRepo struct {
	db *gorm.DB
}

func NewPostgresSeatRepo(db *gorm.DB) *PostgresSeatRepo {
	return &PostgresSeatRepo{db: db}
}

func (r *PostgresSeatRepo) Create(seat *domain.Seat) error {
	return r.db.Create(seat).Error
}

// ฟังก์ชันค้นหาที่นั่งจาก ID
func (r *PostgresSeatRepo) FindByID(id uint) (*domain.Seat, error) {
	var seat domain.Seat
	if err := r.db.First(&seat, id).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}

// ฟังก์ชันอัปเดตสถานะที่นั่ง (เช่น เปลี่ยนจาก AVAILABLE -> LOCKED)
func (r *PostgresSeatRepo) Update(seat *domain.Seat) error {
	return r.db.Save(seat).Error
}

// ฟังก์ชัน Lock ที่นั่ง (กันคนแย่งกัน)
// รับ tx (Transaction) เข้ามา เพราะการ Lock ต้องทำใน Transaction เดียวกันเท่านั้น
func (r *PostgresSeatRepo) GetSeatWithLock(tx *gorm.DB, seatID uint) (*domain.Seat, error) {
	var seat domain.Seat

	// คำสั่ง SQL: SELECT * FROM seats WHERE id = ? FOR UPDATE
	// ใครมาช้า ต้องรอจนกว่า Transaction นี้จะจบ
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&seat, seatID).Error
	return &seat, err
}
