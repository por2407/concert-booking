package domain

import "time"

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique;not null" json:"email"`
	// ใช้ column:password_hash เพื่อ map ให้ตรงกับ database
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	Role         string    `gorm:"default:'USER'" json:"role"` // USER หรือ ADMIN
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
