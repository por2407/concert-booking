package repositories

import (
	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type postgresEventRepo struct {
	db *gorm.DB
}

func NewPostgresEventRepo(db *gorm.DB) *postgresEventRepo {
	return &postgresEventRepo{db: db}
}

func (r *postgresEventRepo) Create(event *domain.Event) error {
	return r.db.Create(event).Error
}
