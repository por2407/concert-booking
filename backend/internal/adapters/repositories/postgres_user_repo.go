package repositories

import (
	"context"

	"github.com/ticket-backend/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(ctx context.Context, user *domain.User) error {
	return GetTx(ctx, r.db).Create(user).Error
}

func (r *PostgresUserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := GetTx(ctx, r.db).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := GetTx(ctx, r.db).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
