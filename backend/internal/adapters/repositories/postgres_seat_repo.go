package repositories

import (
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type postgresSeatRepo struct {
	db *gorm.DB
}

func NewPostgresSeatRepo(db *gorm.DB) *postgresSeatRepo {
	return &postgresSeatRepo{db: db}
}

func (r *postgresSeatRepo) Create(seat *domain.Seat) error {
	return r.db.Create(seat).Error
}
