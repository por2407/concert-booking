package repositories

import (
	"context"

	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresEventRepo struct {
	db *gorm.DB
}

func NewPostgresEventRepo(db *gorm.DB) *PostgresEventRepo {
	return &PostgresEventRepo{db: db}
}

func (r *PostgresEventRepo) Create(ctx context.Context, event *domain.Event) error {
	return GetTx(ctx, r.db).Create(event).Error
}

func (r *PostgresEventRepo) GetByIDWithSeats(ctx context.Context, id uint) (*domain.Event, error) {
	var event domain.Event
	err := GetTx(ctx, r.db).Preload("Seats", func(db *gorm.DB) *gorm.DB {
		return db.Order("row_label ASC, seat_number ASC")
	}).First(&event, id).Error

	return &event, err
}

func (r *PostgresEventRepo) GetAll(ctx context.Context) ([]domain.Event, error) {
	var events []domain.Event
	err := GetTx(ctx, r.db).Find(&events).Error
	return events, err
}
