package repositories

import (
	"context"

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

func (r *PostgresSeatRepo) Create(ctx context.Context, seat *domain.Seat) error {
	return GetTx(ctx, r.db).Create(seat).Error
}

func (r *PostgresSeatRepo) FindByID(ctx context.Context, id uint) (*domain.Seat, error) {
	var seat domain.Seat
	if err := GetTx(ctx, r.db).First(&seat, id).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}

func (r *PostgresSeatRepo) Update(ctx context.Context, seat *domain.Seat) error {
	return GetTx(ctx, r.db).Save(seat).Error
}

func (r *PostgresSeatRepo) GetSeatWithLock(ctx context.Context, seatID uint) (*domain.Seat, error) {
	var seat domain.Seat
	err := GetTx(ctx, r.db).Clauses(clause.Locking{Strength: "UPDATE"}).First(&seat, seatID).Error
	return &seat, err
}
