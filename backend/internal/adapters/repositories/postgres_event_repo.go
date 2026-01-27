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
