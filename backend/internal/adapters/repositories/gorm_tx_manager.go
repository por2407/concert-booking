package repositories

import (
	"context"

	"gorm.io/gorm"
)

type txKey struct{}

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{db: db}
}

// WithTransaction เริ่มต้น Transaction และเก็บ *gorm.DB ไว้ใน Context
func (m *GormTransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// เก็บ tx ไว้ใน context
		ctxWithTx := context.WithValue(ctx, txKey{}, tx)
		return fn(ctxWithTx)
	})
}

// GetTx ดึง *gorm.DB ออกมาจาก Context (ถ้ามี) ถ้าไม่มีให้ใช้ db หลัก
func GetTx(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return defaultDB.WithContext(ctx)
}
