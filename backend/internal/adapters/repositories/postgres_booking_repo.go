package repositories

import (
	"context"

	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresBookingRepo struct {
	db *gorm.DB
}

func NewPostgresBookingRepo(db *gorm.DB) *PostgresBookingRepo {
	return &PostgresBookingRepo{db: db}
}

func (r *PostgresBookingRepo) Create(ctx context.Context, booking *domain.Booking) error {
	return GetTx(ctx, r.db).Create(booking).Error
}

func (r *PostgresBookingRepo) GetByID(ctx context.Context, id uint) (*domain.Booking, error) {
	var booking domain.Booking
	if err := GetTx(ctx, r.db).Preload("Items").First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *PostgresBookingRepo) Update(ctx context.Context, booking *domain.Booking) error {
	return GetTx(ctx, r.db).Save(booking).Error
}

func (r *PostgresBookingRepo) GetByUserID(ctx context.Context, userID uint) ([]domain.Booking, error) {
	var bookings []domain.Booking
	// Preload "Event" เพื่อดูชื่อคอนเสิร์ต
	// Preload "Items.Seat" เพื่อดูว่าจองที่นั่งเบอร์อะไรบ้าง
	err := GetTx(ctx, r.db).
		Preload("Event").
		Preload("Items.Seat").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&bookings).Error

	if err != nil {
		return nil, err
	}
	return bookings, nil
}
